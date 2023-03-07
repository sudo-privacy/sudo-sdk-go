package offlinetask

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/sudo-privacy/sudo-sdk-go/v2/pkg/sudoclient"
	"github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/infra_adapter"
	"github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/task"
	basicvtable "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/vtable"
)

type Project struct {
	ProjectID uint64
	TusitaID  string
	// 本方partyID
	PartyID string
}

func (project Project) Valid() error {
	if project.ProjectID == 0 {
		return errors.New("projectID not set")
	}
	if project.TusitaID == "" {
		return errors.New("tusita not set")
	}
	if project.PartyID == "" {
		return errors.New("self party not set")
	}
	return nil
}

type Factory struct {
	*sudoclient.SudoClient
	Project
}

func NewFactory(
	client *sudoclient.SudoClient,
	project Project) (*Factory, error) {
	if err := project.Valid(); err != nil {
		return nil, err
	}
	return &Factory{
		SudoClient: client,
		Project:    project,
	}, nil
}

// OperatorInput 配置Dag中多个算子的输入
type OperatorInput struct {
	Stages map[string]StageInput
}

// StageInput 配置单个算子中多个参与方的输入
type StageInput struct {
	Client  PartyInput
	Servers []PartyInput
}

// PartyInput 配置一个算子中的一个参与方的一种输入
type PartyInput struct {
	PartyID string
	Vtables map[string]VTableInput
}

// VTableInput 配置Vtable输入
type VTableInput struct {
	VtableID uint64
	// 索引列
	KeyColumn          string
	SelectedVtableCols []string
}

func (input *OperatorInput) GetInputVtable(stage string, innerID uint64, party, vtable string) (VTableInput, error) {
	empty := VTableInput{}
	if _, ok := input.Stages[stage]; !ok {
		return empty, errors.Errorf("stage %s input not found", stage)
	}
	parties := []PartyInput{}
	parties = append(parties, input.Stages[stage].Client)
	parties = append(parties, input.Stages[stage].Servers...)
	if len(parties)-1 >= int(innerID) {
		if parties[innerID].PartyID == party {
			if vtable, ok := parties[innerID].Vtables[vtable]; ok {
				return vtable, nil
			}
		}
	}
	return empty, errors.Errorf("stage %s party %s input %s not found", stage, party, vtable)
}

type OperatorParams struct {
	Stages map[string]StageParams
}

type StageParams struct {
	Client       StagePartyParams
	Servers      []StagePartyParams
	SharedParams map[string]interface{}
}

type StagePartyParams struct {
	PartyID string
	Params  map[string]interface{}
}

func (params *OperatorParams) GetPartyParams(stage string, innerID uint64, party, param string) interface{} {
	if _, ok := params.Stages[stage]; !ok {
		return nil
	}
	parties := []StagePartyParams{}
	parties = append(parties, params.Stages[stage].Client)
	parties = append(parties, params.Stages[stage].Servers...)
	if len(parties)-1 >= int(innerID) {
		if parties[innerID].PartyID == party {
			if iParam, ok := parties[innerID].Params[param]; ok {
				return iParam
			}
		}
	}
	return nil
}

func (params *OperatorParams) GetShareParams(stage, param string) interface{} {
	if _, ok := params.Stages[stage]; !ok {
		return nil
	}
	if iParam, ok := params.Stages[stage].SharedParams[param]; ok {
		return iParam
	}
	return nil
}

// GetTaskModelOutputs 获取任务各算子每个参与方的 .model类型vtable输出
func (f *Factory) GetTaskModelOutputs(ctx context.Context, taskID uint64) (*OperatorInput, error) {
	taskInfo, err := f.SudoClient.GetTasks(ctx, &task.GetTasksRequest{
		TaskId:        taskID,
		ToTusita:      f.TusitaID,
		WithJobConfig: true,
	})
	resp := OperatorInput{
		Stages: make(map[string]StageInput),
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
		stageInput := StageInput{
			Servers: make([]PartyInput, 0),
		}
		partiesConfigs := []*task.SiteConfig{}
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ClientConfig)
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ServerConfigs...)
		for j := range partiesConfigs {
			if partiesConfigs[j] == nil {
				continue
			}
			partyVtableInput := PartyInput{
				Vtables: make(map[string]VTableInput),
			}
			for k := range partiesConfigs[j].Outputs {
				if strings.Contains(k, ".model") || strings.Contains(k, ".smodel") || strings.Contains(k, ".cmodel") {
					vtable, getOutputErr := f.GetStageOutputVtable(
						ctx, taskModel.Stages[i].JobId,
						partiesConfigs[j].PartyId,
						uint64(j),
						k)
					if getOutputErr != nil {
						return nil, getOutputErr
					}
					partyVtableInput.PartyID = partiesConfigs[j].PartyId
					partyVtableInput.Vtables[k] = VTableInput{VtableID: vtable.Base.Id, SelectedVtableCols: []string{"*"}}
				}
			}
			if j == 0 {
				stageInput.Client = partyVtableInput
			} else {
				stageInput.Servers = append(stageInput.Servers, partyVtableInput)
			}
		}
		resp.Stages[taskModel.Stages[i].StageName] = stageInput
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (f *Factory) GetStageOutputVtable(
	ctx context.Context,
	stageID uint64,
	partyID string,
	innerID uint64,
	vtable string,
) (*basicvtable.Vtable, error) {
	resp, err := f.SudoClient.GetStageOutputs(ctx, &task.GetStageOutputRequest{
		StageId:  stageID,
		PartyId:  partyID,
		InnerId:  innerID,
		ToTusita: f.TusitaID,
	})
	if err != nil {
		return nil, errors.Wrapf(err,
			"f.SudoClient.GetStageOutputs stage %d,partyID %s",
			stageID, partyID)
	}
	for vtableK, vtableV := range resp.Vtables {
		if strings.Contains(vtableK, vtable) {
			return vtableV, nil
		}
	}

	return nil, errors.Errorf("stage %d, party %s output %s not found",
		stageID, partyID, vtable)
}

// SetPredictTaskModelInputFromOutput 给任务配置 模型和数据 vtable数据输入。
func (f *Factory) SetTaskInput(taskModel *task.Task, modelInput, dataInput *OperatorInput) (*task.Task, error) {
	if taskModel == nil {
		return nil, errors.New("task is nil")
	}
	for i := range taskModel.Stages {
		if taskModel.Stages[i].Config == nil {
			continue
		}
		stageName := taskModel.Stages[i].StageName
		partiesConfigs := []*task.SiteConfig{}
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ClientConfig)
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ServerConfigs...)
		for j := range partiesConfigs {
			if partiesConfigs[j] == nil {
				continue
			}
			for k := range partiesConfigs[j].Inputs {
				if partiesConfigs[j].Inputs[k].Type != "vtable" {
					continue
				}
				if strings.Contains(k, ".model") || strings.Contains(k, ".smodel") || strings.Contains(k, ".cmodel") {
					partyID := partiesConfigs[j].PartyId
					vtable, err := modelInput.GetInputVtable(stageName, uint64(j), partyID, k)
					if err != nil {
						return nil, err
					}
					partiesConfigs[j].Inputs[k] = &infra_adapter.FurnaceProperty{
						Id:                 vtable.VtableID,
						Type:               "vtable",
						SelectedVtableCols: vtable.SelectedVtableCols,
					}
				}
				if strings.Contains(k, ".table") {
					partyID := partiesConfigs[j].PartyId
					vtable, err := dataInput.GetInputVtable(stageName, uint64(j), partyID, k)
					if err != nil {
						return nil, err
					}
					partiesConfigs[j].Inputs[k] = &infra_adapter.FurnaceProperty{
						Id:                 vtable.VtableID,
						Type:               "vtable",
						SelectedVtableCols: vtable.SelectedVtableCols,
					}
				}
			}
		}
	}
	return taskModel, nil
}

// setTaskOutputDataSrc 设置各个参与方所有算子输出数据源。
func (f *Factory) SetTaskOutputDataSrc(taskModel *task.Task, outputDataSrc map[string]uint64) (*task.Task, error) {
	if taskModel == nil {
		return nil, errors.New("task is nil")
	}
	for i := range taskModel.Stages {
		if taskModel.Stages[i].Config == nil {
			continue
		}
		partiesConfigs := []*task.SiteConfig{}
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ClientConfig)
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ServerConfigs...)
		for j := range partiesConfigs {
			if partiesConfigs[j] == nil {
				continue
			}
			for k := range partiesConfigs[j].Outputs {
				if _, ok := outputDataSrc[partiesConfigs[j].PartyId]; !ok {
					return nil, errors.Errorf("party %s output datasrc not set", partiesConfigs[j].PartyId)
				}
				partiesConfigs[j].Outputs[k] = &infra_adapter.FurnaceProperty{
					Id:   outputDataSrc[partiesConfigs[j].PartyId],
					Type: "datasrc",
				}
			}
		}
	}
	return taskModel, nil
}

// setTaskOutputDataSrc 设置各个参与方所有算子输出数据源。
func (f *Factory) ReplaceParty(taskModel *task.Task, partyMap map[string]string) (*task.Task, error) {
	if taskModel == nil {
		return nil, errors.New("task is nil")
	}
	for i := range taskModel.Stages {
		if taskModel.Stages[i].Config == nil {
			continue
		}
		partiesConfigs := []*task.SiteConfig{}
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ClientConfig)
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ServerConfigs...)
		for j := range partiesConfigs {
			if partiesConfigs[j] == nil {
				continue
			}
			newParty, ok := partyMap[partiesConfigs[j].PartyId]
			if !ok {
				return nil, errors.Errorf("party %s not set", partiesConfigs[j].PartyId)
			}
			partiesConfigs[j].PartyId = newParty
		}
	}
	return taskModel, nil
}

func (f *Factory) SetTaskOperatorParams(taskModel *task.Task, params OperatorParams) (*task.Task, error) {
	if taskModel == nil {
		return nil, errors.New("task is nil")
	}
	for i := range taskModel.Stages {
		if taskModel.Stages[i].Config == nil {
			continue
		}
		stageName := taskModel.Stages[i].StageName
		partiesConfigs := []*task.SiteConfig{}
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ClientConfig)
		partiesConfigs = append(partiesConfigs, taskModel.Stages[i].Config.ServerConfigs...)
		for j := range partiesConfigs {
			if partiesConfigs[j] == nil {
				continue
			}
			for k := range partiesConfigs[j].Params {
				partyID := partiesConfigs[j].PartyId
				iParam := params.GetPartyParams(stageName, uint64(j), partyID, k)
				if iParam != nil {
					structpbValue, err := structpb.NewValue(iParam)
					if err != nil {
						return nil, errors.Wrapf(err, "structpb.NewValue(iParam) failed %+v", iParam)
					}
					partiesConfigs[j].Params[k] = structpbValue
				}
			}
		}
		for k := range taskModel.Stages[i].Config.SharedParams {
			iParam := params.GetShareParams(stageName, k)
			if iParam != nil {
				structpbValue, err := structpb.NewValue(iParam)
				if err != nil {
					return nil, errors.Wrapf(err, "structpb.NewValue(iParam) failed %+v", iParam)
				}
				taskModel.Stages[i].Config.SharedParams[k] = structpbValue
			}
		}
	}
	return taskModel, nil
}
