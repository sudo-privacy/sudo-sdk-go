package onlinesvc

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

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
	projectID uint64
	tusitaID  string
	// 本方partyID
	partyID string
}

func NewFactory(
	client *sudoclient.SudoClient,
	tusitaID string,
	projectID uint64,
	partyID string) *Factory {
	return &Factory{
		SudoClient: client,
		projectID:  projectID,
		tusitaID:   tusitaID,
		partyID:    partyID,
	}
}

// OperatorVtableInput 配置Dag中多个算子的输入
type OperatorVtableInput struct {
	Stage map[string]StageVtableInput
}

// StageVtableInput 配置单个算子中多个参与方的输入
type StageVtableInput struct {
	Input map[string]PartyVtableInput
}

// PartyVtableInput 配置一个算子中的一个参与方的一种输入
type PartyVtableInput struct {
	Input map[string]VTableInput
}

// VTableInput 配置Vtable输入
type VTableInput struct {
	VtableID uint64
	// 索引列
	KeyColumn          string
	selectedVtableCols []string
}

// 检查预测任务使用合法性
func (f *Factory) verifyPredictTask(ctx context.Context, taskID uint64) error {
	taskInfo, err := f.SudoClient.GetTasks(ctx, &task.GetTasksRequest{
		TaskId:    taskID,
		ToTusita:  f.tusitaID,
		ProjectId: f.projectID,
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

	if taskInfo.Data[0].Base.Base.OwnerParty != f.partyID {
		return errors.New("task owner is not this party")
	}

	status := taskInfo.Data[0].Base.Base.Status
	if status != enums.Task_SUCCESS && status != enums.Task_USING {
		return errors.New("task is not in SUCCESS or USING status")
	}
	return nil
}

// 检查onlineBlock input配置和预测任务算子的.table输入配置对齐
func (f *Factory) verifyOnlineInputs(ctx context.Context, taskID uint64, onlineInput *OperatorVtableInput) error {
	taskInputs, err := f.GetTaskVTableInputs(ctx, taskID)
	if err != nil {
		return err
	}
	err = f.verifyTables(ctx, onlineInput)
	if err != nil {
		return err
	}
	for stage, stageInput := range taskInputs.Stage {
		_, ok := onlineInput.Stage[stage]
		if !ok {
			return errors.Errorf("stage %s input not found", stage)
		}
		for party, partyInputs := range stageInput.Input {
			_, ok := onlineInput.Stage[stage].Input[party]
			if !ok {
				return errors.Errorf("stage %s party %s input not found", stage, party)
			}
			for inputKey := range partyInputs.Input {
				_, ok := onlineInput.Stage[stage].Input[party].Input[inputKey]
				if !ok {
					return errors.Errorf("stage %s party %s input %s not found", stage, party, inputKey)
				}
			}
		}
	}
	return nil
}

// GetTaskVTableInputs 收集预测任务中所有算子的 .table 类型输入，可用于引导onlineBlock Input配置。
func (f *Factory) GetTaskVTableInputs(ctx context.Context, taskID uint64) (*OperatorVtableInput, error) {
	taskInfo, err := f.SudoClient.GetTasks(ctx, &task.GetTasksRequest{
		TaskId:        taskID,
		ToTusita:      f.tusitaID,
		WithJobConfig: true,
	})
	resp := OperatorVtableInput{
		Stage: make(map[string]StageVtableInput),
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
	task := taskInfo.Data[0].Base.Base
	for i := range task.Stages {
		if task.Stages[i] == nil || task.Stages[i].Config == nil {
			continue
		}
		stageVtableInput := StageVtableInput{
			Input: make(map[string]PartyVtableInput),
		}

		if task.Stages[i].Config.ClientConfig != nil {
			partyVtableInput := PartyVtableInput{
				Input: make(map[string]VTableInput),
			}
			for k, v := range task.Stages[i].Config.ClientConfig.Inputs {
				if v.Type == "vtable" && strings.Contains(k, ".table") {
					partyVtableInput.Input[k] = VTableInput{VtableID: v.Id, selectedVtableCols: v.SelectedVtableCols}
				}
			}
			if len(partyVtableInput.Input) > 0 {
				stageVtableInput.Input[task.Stages[i].Config.ClientConfig.PartyId] = partyVtableInput
			}
		}
		for j := range task.Stages[i].Config.ServerConfigs {
			if task.Stages[i].Config.ServerConfigs[j] == nil {
				continue
			}
			partyVtableInput := PartyVtableInput{
				Input: make(map[string]VTableInput),
			}
			for k, v := range task.Stages[i].Config.ServerConfigs[j].Inputs {
				if v.Type == "vtable" && strings.Contains(k, ".table") {
					partyVtableInput.Input[k] = VTableInput{VtableID: v.Id, selectedVtableCols: v.SelectedVtableCols}
				}
			}
			if len(partyVtableInput.Input) > 0 {
				stageVtableInput.Input[task.Stages[i].Config.ServerConfigs[j].PartyId] = partyVtableInput
			}
		}
		if len(stageVtableInput.Input) > 0 {
			resp.Stage[task.Stages[i].StageName] = stageVtableInput
		}
	}
	err = f.fetchInputsColumns(ctx, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 补充输入table的具体column配置，目前未使用
func (f *Factory) fetchInputsColumns(ctx context.Context, inputs *OperatorVtableInput) error {
	for stage := range inputs.Stage {
		for party := range inputs.Stage[stage].Input {
			for inputKey, vtableInput := range inputs.Stage[stage].Input[party].Input {
				if len(vtableInput.selectedVtableCols) == 0 {
					vtableInput.selectedVtableCols = []string{"*"}
				}
				if len(vtableInput.selectedVtableCols) == 1 && vtableInput.selectedVtableCols[0] == "*" {
					// add to-tusita metadata
					newCtx := metadata.AppendToOutgoingContext(ctx, "to-tusita", f.tusitaID)
					resp, err := f.SudoClient.ListVtables(newCtx, &vtable.ListVtablesRequest{
						QueryOptions: &basicvtable.VtableQueryOptions{
							Id:               vtableInput.VtableID,
							WithSegmentation: true,
							PartyId:          party,
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
					vtableInput.selectedVtableCols = columns
				}
				inputs.Stage[stage].Input[party].Input[inputKey] = vtableInput
			}
		}
	}
	return nil
}

// 检查vtable在project中可见，且为MYSQL类型
func (f *Factory) verifyTables(ctx context.Context, inputs *OperatorVtableInput) error {
	for stage := range inputs.Stage {
		for party := range inputs.Stage[stage].Input {
			for _, vtableInput := range inputs.Stage[stage].Input[party].Input {
				// add to-tusita metadata
				newCtx := metadata.AppendToOutgoingContext(ctx, "to-tusita", f.tusitaID)
				resp, err := f.SudoClient.ListVtables(newCtx, &vtable.ListVtablesRequest{
					QueryOptions: &basicvtable.VtableQueryOptions{
						Id:               vtableInput.VtableID,
						WithSegmentation: true,
						PartyId:          party,
						WithDatasrc:      true,
						WithAdditional:   true,
						ProjectId:        f.projectID,
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
	inputs *OperatorVtableInput,
) (*online_service.OnlineBlockTemplate, error) {
	resp, err := f.SudoClient.GetOnlineBlockFromModel(ctx, &online_service.GetOnlineBlockFromModelRequest{
		TaskId:       taskID,
		MajorVersion: 1,
		MinorVersion: 0,
		ToTusita:     f.tusitaID,
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

// 将input配置到OnlineBlock模板中
func (f *Factory) setOnlineBlockTableInputs(
	onlineBlock *online_service.OnlineBlockTemplate,
	inputs *OperatorVtableInput,
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
					stageInput, ok := inputs.Stage[onlineBlock.Operators[i].Name]
					if !ok {
						return nil, errors.Errorf("stage %s input not found", onlineBlock.Operators[i].Name)
					}
					partyInput, ok := stageInput.Input[onlineBlock.Operators[i].Participants[j].PartyId]
					if !ok {
						return nil, errors.Errorf("stage %s party %s input not found",
							onlineBlock.Operators[i].Name,
							onlineBlock.Operators[i].Participants[j].PartyId,
						)
					}
					vtableInput, ok := partyInput.Input[k]
					if !ok {
						return nil, errors.Errorf("stage %s party %s input %s not found",
							onlineBlock.Operators[i].Name,
							onlineBlock.Operators[i].Participants[j].PartyId,
							k,
						)
					}
					onlineBlock.Operators[i].Participants[j].OnlineInputs[k] = &online_service.OnlineInput{
						VtableId:  vtableInput.VtableID,
						KeyColumn: vtableInput.KeyColumn,
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
	onlineInputs *OperatorVtableInput) (*Service, error) {
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
			ProjectId:    f.projectID,
			Path:         sudoclient.RandPath(6),
			ProtectId:    protectID,
			TaskId:       taskID,
			MajorVersion: 1,
			MinorVersion: 0,
			OnlineBlock:  onlineBlock,
		},
		ToTusita: f.tusitaID,
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
	onlineInputs *OperatorVtableInput) (*Service, error) {
	resp, err := f.SudoClient.GetServices(ctx, &online_service.GetServicesRequest{
		Filter: &online_service.ServiceFilter{
			Pager: &paginator.Paginator{
				NotPaging: true,
			},
			OwnerPartyId: f.partyID,
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
	query.ToTusita = f.tusitaID
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
		ToTusita:  f.tusitaID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get service failed")
	}
	return &Service{
		Factory: f,
		Service: resp.Data.Service,
	}, nil
}
