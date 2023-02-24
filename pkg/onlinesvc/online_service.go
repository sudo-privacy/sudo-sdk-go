package onlinesvc

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	"github.com/sudo-privacy/sudo-sdk-go/pkg/offlinetask"
	"github.com/sudo-privacy/sudo-sdk-go/pkg/sudoclient"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/online_service"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/paginator"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/task"
	basicvtable "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/vtable"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/vtable"
)

// Factory 是对 [sudoclient.SudoClient] 的简单封装，其中配置了中心节点和项目，提供创建在线服务方法。
type Factory struct {
	*sudoclient.SudoClient
	offlinetask.Project
}

func NewFactory(
	client *sudoclient.SudoClient,
	project offlinetask.Project) (*Factory, error) {
	if err := project.Valid(); err != nil {
		return nil, err
	}
	return &Factory{
		SudoClient: client,
		Project:    project,
	}, nil
}

// 检查预测任务使用合法性
func (f *Factory) verifyPredictTask(ctx context.Context, taskID uint64) error {
	taskInfo, err := f.SudoClient.GetTasks(ctx, &task.GetTasksRequest{
		TaskId:    taskID,
		ToTusita:  f.TusitaID,
		ProjectId: f.ProjectID,
	})
	if err != nil {
		return errors.Wrap(err, "get tasks error")
	}
	if len(taskInfo.Data) != 1 {
		return errors.New("query single task,return task length != 1")
	}
	if taskInfo.Data[0].Base == nil || taskInfo.Data[0].Base.Base == nil {
		return errors.New("task field is nil")
	}

	if taskInfo.Data[0].Base.Base.OwnerParty != f.PartyID {
		return errors.New("task owner is not this party")
	}

	status := taskInfo.Data[0].Base.Base.Status
	if status != enums.Task_SUCCESS && status != enums.Task_USING {
		return errors.New("task is not in SUCCESS or USING status")
	}
	return nil
}

// 检查onlineBlock input配置和预测任务算子的.table输入配置对齐
func (f *Factory) verifyOnlineInputs(ctx context.Context, taskID uint64, onlineInput *offlinetask.OperatorInput) error {
	taskInputs, err := f.GetTaskVTableInputs(ctx, taskID)
	if err != nil {
		return err
	}
	err = f.verifyTables(ctx, onlineInput)
	if err != nil {
		return err
	}
	for stage, stageInput := range taskInputs.Stages {
		partiesInput := []offlinetask.PartyInput{}
		partiesInput = append(partiesInput, stageInput.Client)
		partiesInput = append(partiesInput, stageInput.Servers...)
		for i := range partiesInput {
			for j := range partiesInput[i].Vtables {
				_, err = onlineInput.GetInputVtable(stage, uint64(i), partiesInput[i].PartyID, j)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// GetTaskVTableInputs 收集预测任务中所有算子的 .table 类型输入，可用于引导onlineBlock Input配置。
func (f *Factory) GetTaskVTableInputs(ctx context.Context, taskID uint64) (*offlinetask.OperatorInput, error) {
	taskInfo, err := f.SudoClient.GetTasks(ctx, &task.GetTasksRequest{
		TaskId:        taskID,
		ToTusita:      f.TusitaID,
		WithJobConfig: true,
	})
	resp := offlinetask.OperatorInput{
		Stages: make(map[string]offlinetask.StageInput),
	}
	if err != nil {
		return nil, errors.Wrap(err, "get tasks error")
	}
	if len(taskInfo.Data) != 1 {
		return nil, errors.New("query single task,return task length != 1")
	}
	if taskInfo.Data[0].Base == nil || taskInfo.Data[0].Base.Base == nil {
		return nil, errors.New("task field is nil")
	}
	taskModel := taskInfo.Data[0].Base.Base
	for i := range taskModel.Stages {
		if taskModel.Stages[i] == nil || taskModel.Stages[i].Config == nil {
			continue
		}
		stageInput := offlinetask.StageInput{
			Servers: make([]offlinetask.PartyInput, 0),
		}
		partiesConfigs := []*task.SiteConfig{}
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ClientConfig)
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ServerConfigs...)

		for j := range partiesConfigs {
			if partiesConfigs[j] == nil {
				continue
			}
			partyVtableInput := offlinetask.PartyInput{
				PartyID: partiesConfigs[j].PartyId,
				Vtables: make(map[string]offlinetask.VTableInput),
			}
			for k, v := range partiesConfigs[j].Inputs {
				if v.Type == "vtable" && strings.Contains(k, ".table") {
					partyVtableInput.Vtables[k] = offlinetask.VTableInput{VtableID: v.Id, SelectedVtableCols: v.SelectedVtableCols}
				}
			}
			if len(partyVtableInput.Vtables) > 0 {
				if j == 0 {
					stageInput.Client = partyVtableInput
				} else {
					stageInput.Servers = append(stageInput.Servers, partyVtableInput)
				}
			}
		}
		resp.Stages[taskModel.Stages[i].StageName] = stageInput
	}
	err = f.fetchInputsColumns(ctx, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 补充输入table的具体column配置，目前未使用
func (f *Factory) fetchInputsColumns(ctx context.Context, inputs *offlinetask.OperatorInput) error {
	for stage := range inputs.Stages {
		parties := []offlinetask.PartyInput{}
		parties = append(parties, inputs.Stages[stage].Client)
		parties = append(parties, inputs.Stages[stage].Servers...)
		for i := range parties {
			for inputKey, vtableInput := range parties[i].Vtables {
				if len(vtableInput.SelectedVtableCols) == 0 {
					vtableInput.SelectedVtableCols = []string{"*"}
				}
				if len(vtableInput.SelectedVtableCols) == 1 && vtableInput.SelectedVtableCols[0] == "*" {
					// add to-tusita metadata
					newCtx := metadata.AppendToOutgoingContext(ctx, "to-tusita", f.TusitaID)
					resp, err := f.SudoClient.ListVtables(newCtx, &vtable.ListVtablesRequest{
						QueryOptions: &basicvtable.VtableQueryOptions{
							Id:               vtableInput.VtableID,
							WithSegmentation: true,
							PartyId:          parties[i].PartyID,
						},
					})
					if err != nil {
						return errors.Wrap(err, "list vtable failed")
					}
					if len(resp.Data) != 1 {
						return errors.New("get vtable length != 1")
					}
					if resp.Data[0].Base == nil {
						return errors.New("vtable data is nil")
					}
					columns := []string{}
					for i := range resp.Data[0].Base.Columns {
						columns = append(columns, resp.Data[0].Base.Columns[i].Field)
					}
					vtableInput.SelectedVtableCols = columns
				}
				if i == 0 {
					inputs.Stages[stage].Client.Vtables[inputKey] = vtableInput
				} else {
					inputs.Stages[stage].Servers[i-1].Vtables[inputKey] = vtableInput
				}
			}
		}
	}
	return nil
}

// 检查vtable在project中可见，且为MYSQL类型
func (f *Factory) verifyTables(ctx context.Context, inputs *offlinetask.OperatorInput) error {
	for stage := range inputs.Stages {
		parties := []offlinetask.PartyInput{}
		parties = append(parties, inputs.Stages[stage].Client)
		parties = append(parties, inputs.Stages[stage].Servers...)
		for i := range parties {
			for _, vtableInput := range parties[i].Vtables {
				// add to-tusita metadata
				newCtx := metadata.AppendToOutgoingContext(ctx, "to-tusita", f.TusitaID)
				resp, err := f.SudoClient.ListVtables(newCtx, &vtable.ListVtablesRequest{
					QueryOptions: &basicvtable.VtableQueryOptions{
						Id:               vtableInput.VtableID,
						WithSegmentation: true,
						PartyId:          parties[i].PartyID,
						WithDatasrc:      true,
						WithAdditional:   true,
						ProjectId:        f.ProjectID,
					},
				})
				if err != nil {
					return errors.Wrap(err, "list vtable failed")
				}
				if len(resp.Data) != 1 {
					return errors.New("get vtable length != 1")
				}
				if resp.Data[0].Base == nil {
					return errors.New("vtable data is nil")
				}
				if resp.Data[0].Datasrc == nil {
					return errors.New("vtable dataSrc field is nil")
				}
				if resp.Data[0].Datasrc.Type != enums.Datasource_MYSQL {
					return errors.Errorf("only support Datasource_MYSQL vtable, vtable %d is %s",
						vtableInput.VtableID,
						resp.Data[0].Datasrc.Type.String(),
					)
				}
			}
		}
	}
	return nil
}

// 1.从预测任务转化 onlineBlock模板，
// 2.使用input配置 onlineBlock模板中的onlineInput
func (f *Factory) getOnlineBlock(
	ctx context.Context, taskID uint64,
	inputs *offlinetask.OperatorInput,
) (*online_service.OnlineBlockTemplate, error) {
	resp, err := f.SudoClient.GetOnlineBlockFromModel(ctx, &online_service.GetOnlineBlockFromModelRequest{
		TaskId:       taskID,
		MajorVersion: 1,
		MinorVersion: 0,
		ToTusita:     f.TusitaID,
	})
	if err != nil {
		return nil, err
	}
	onlineBlock, err := f.setOnlineBlockTableInputs(resp.Data, inputs)
	if err != nil {
		return nil, err
	}

	return onlineBlock, nil
}

// setOnlineBlockTableInputs 将input配置到OnlineBlock模板中
func (f *Factory) setOnlineBlockTableInputs(
	onlineBlock *online_service.OnlineBlockTemplate,
	inputs *offlinetask.OperatorInput,
) (*online_service.OnlineBlockTemplate, error) {
	for i := range onlineBlock.Operators {
		if onlineBlock.Operators[i] == nil {
			continue
		}
		for j := range onlineBlock.Operators[i].Participants {
			if onlineBlock.Operators[i].Participants[j] == nil {
				continue
			}
			for k := range onlineBlock.Operators[i].Participants[j].OnlineInputs {
				if onlineBlock.Operators[i].Participants[j].OnlineInputs[k] == nil {
					continue
				}
				if strings.Contains(k, ".table") {
					vtable, err := inputs.GetInputVtable(
						onlineBlock.Operators[i].Name, uint64(j),
						onlineBlock.Operators[i].Participants[j].PartyId,
						k)
					if err != nil {
						return nil, err
					}
					onlineBlock.Operators[i].Participants[j].OnlineInputs[k] = &online_service.OnlineInput{
						VtableId:  vtable.VtableID,
						KeyColumn: vtable.KeyColumn,
					}
				}
			}
		}
	}
	return onlineBlock, nil
}

// CreateService 创建在线服务。在线服务需要多个参与方支持，本方发起创建服务请求，创建后服务的OwnerParty也就是本方。
// 服务创建后需要其他参与方审批后服务才能真正部署并使用。
// 使用预测任务和输入配置作为输入，预测任务会被转化onlineBlockTemplate，输入配置则配置到onlineblockTemplate中的onlineInput中。最后onlineBlockTemplate被部署为在线服务。
// 在线服务中的onlineInput只允许MySQL类型的vtable,因为在线服务提供预测服务服务时需要从onlineInput配置的vtable中查询特征，综合多个参与方的特征列后使用模型完成预测。
// 对特征列的查询需要使用索引，所以限制只允许MySQL类型的vtable输入（后期可能重构后移除这个限制）。
//
//	- name 指定在线服务名称
//	- protectID 在线服务是否保护id，保护id会使查询参与方特征列时使用pir。
//	- taskID 预测任务ID,该预测任务必须属于当前项目，且必须是成功执行的任务。
//	- onlineInputs,在线服务输入，必须和预测任务生成的onlineBlockTemplate中的onlineInput对齐。
//	  可以使用 [Factory.GetTaskVTableInputs] 获取输入配置的格式，并参考选择vtable。
//	  vtable需要当前项目可见且为MySQL类型，不限vtable必须包含特征列，也不限制多个参与方的索引必须匹配。缺失的特征列会被0填充并影响预测结果准确性。
func (f *Factory) CreateService(
	ctx context.Context,
	name string,
	protectID bool,
	taskID uint64,
	onlineInputs *offlinetask.OperatorInput) (*Service, error) {
	err := f.verifyPredictTask(ctx, taskID)
	if err != nil {
		return nil, err
	}
	err = f.verifyOnlineInputs(ctx, taskID, onlineInputs)
	if err != nil {
		return nil, err
	}
	onlineBlock, err := f.getOnlineBlock(ctx, taskID, onlineInputs)
	if err != nil {
		return nil, err
	}

	resp, err := f.SudoClient.CreateService(ctx, &online_service.CreateServiceRequest{
		Service: &online_service.ServiceParam{
			Name:         name,
			ServiceType:  enums.Service_PREDICT,
			ProjectId:    f.ProjectID,
			Path:         sudoclient.RandPath(6),
			ProtectId:    protectID,
			TaskId:       taskID,
			MajorVersion: 1,
			MinorVersion: 0,
			OnlineBlock:  onlineBlock,
		},
		ToTusita: f.TusitaID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "createService failed")
	}
	return &Service{
		Factory: f,
		Service: resp.Data,
	}, nil
}

// NewService 返回或创建 [Service] 如果名称为 name 的在线服务已经存在，直接返回该在线服务，否则使用参数创建在线服务。
// 参数配置参考 [Factory.CreateService] 。
func (f *Factory) NewService(
	ctx context.Context,
	name string,
	protectID bool,
	taskID uint64,
	onlineInputs *offlinetask.OperatorInput) (*Service, error) {
	resp, err := f.SudoClient.GetServices(ctx, &online_service.GetServicesRequest{
		Filter: &online_service.ServiceFilter{
			Pager: &paginator.Paginator{
				NotPaging: true,
			},
			OwnerPartyId: f.PartyID,
			ServiceType:  "PREDICT",
		},
		Tab: "SELF",
	})
	if err != nil {
		return nil, errors.Wrap(err, "get Services failed")
	}
	for i := range resp.Data {
		if resp.Data[i].Base != nil && resp.Data[i].Base.Name == name {
			return &Service{
				Factory: f,
				Service: resp.Data[i],
			}, nil
		}
	}
	return f.CreateService(ctx, name, protectID, taskID, onlineInputs)
}

// GetServices 查询 [Service] ，根据输入的过滤参数，可能返回0个、1个 或多个 Server。
// tusita参数会从Factory中配置
func (f *Factory) GetServices(ctx context.Context, query *online_service.GetServicesRequest) ([]*Service, error) {
	query.ToTusita = f.TusitaID
	servicesResp, err := f.SudoClient.GetServices(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get servicess")
	}
	resp := []*Service{}
	for i := range servicesResp.Data {
		resp = append(resp, &Service{
			Factory: f,
			Service: servicesResp.Data[i],
		})
	}
	return resp, nil
}

// GetService 使用serviceID 获取 [Service]。
func (f *Factory) GetService(ctx context.Context, serviceID uint64) (*Service, error) {
	resp, err := f.SudoClient.GetService(ctx, &online_service.GetServiceRequest{
		ServiceId: strconv.FormatUint(serviceID, 10),
		ToTusita:  f.TusitaID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get service failed")
	}
	return &Service{
		Factory: f,
		Service: resp.Data.Service,
	}, nil
}
