package onlinesvc

import (
	"context"
	"io"
	"os"

	"github.com/sudo-privacy/sudo-sdk-go/pkg/offlinetask"
	"github.com/sudo-privacy/sudo-sdk-go/pkg/sudoclient"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/online_service"
	prototask "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/task"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/task"
)

// 从预测任务完成后开始的流程
func ExampleService_Predict() {
	sudoClientF1, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("10.0.23.130:8071"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("http://10.0.23.130:7857"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	f1, _ := NewFactory(
		sudoClientF1,
		offlinetask.Project{
			TusitaID:  "tusita-t1",
			ProjectID: 442450667450138891,
			PartyID:   "party-f1",
		})

	sudoClientF2, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("10.0.23.130:18071"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("10.0.23.130:17857"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	f2, _ := NewFactory(sudoClientF2, offlinetask.Project{
		TusitaID:  "tusita-t1",
		ProjectID: 442450667450138891,
		PartyID:   "party-f2",
	})
	// vtable可以使用数牍隐私计算平台上已经存在的，也可以通过 sudoclient.CreateVtableFromLocalFile、sudoclient.CreateVtableFromDB 提前创建。
	input := &offlinetask.OperatorInput{
		Stages: map[string]offlinetask.StageInput{
			"数据预处理（训练）-1": {
				Client: offlinetask.PartyInput{
					PartyID: "party-f1",
					Vtables: map[string]offlinetask.VTableInput{
						"data.table": offlinetask.VTableInput{
							VtableID:  62,
							KeyColumn: "ID",
						},
					},
				},
			},

			"数据预处理（训练）-2": {
				Client: offlinetask.PartyInput{
					PartyID: "party-f2",
					Vtables: map[string]offlinetask.VTableInput{
						"data.table": offlinetask.VTableInput{
							VtableID:  35,
							KeyColumn: "ID",
						},
					},
				},
			},
		},
	}
	// 创建 service
	service, err := f1.NewService(context.TODO(), "lyy-sdk-t-est", false, 18, input)
	if err != nil {
		// TODO: Handle error.
	}
	// 数据提供方审核
	serverService, err := f2.GetService(context.Background(), service.ID())
	if err != nil {
		// TODO: Handle error.
	}
	err = serverService.Approve(context.Background())
	if err != nil {
		// TODO: Handle error.
	}
	// 服务发起方等待服务部署完成
	err = service.WaitForReady(context.Background())
	if err != nil {
		// TODO: Handle error.
	}
	// 预测查询
	_, err = service.Predict(context.Background(), &online_service.PredictRequest{
		Ids: []string{"223"},
	})

	if err != nil {
		// TODO: Handle error.
	}
	err = service.Delete(context.Background())
	if err != nil {
		// TODO: Handle error.
	}
	f1.Close()
	f2.Close()
}

// 从训练任务开始的完成流程
func ExampleService_Predict_all() {
	// 替换模板中的参与方
	partyMap := map[string]string{
		"party-fa": "party-f1",
		"party-fb": "party-f2",
	}
	//替换模板中个参与方的输出数据源
	partyOutputDataSrc := map[string]uint64{
		"party-f1": 2,
		"party-f2": 2,
	}

	// 替换训练任务模板中的参与方和输入
	// vtable可以使用数牍隐私计算平台上已经存在的，也可以通过 sudoclient.CreateVtableFromLocalFile、sudoclient.CreateVtableFromDB 提前创建。
	trainTaskVtableInput := &offlinetask.OperatorInput{
		Stages: map[string]offlinetask.StageInput{
			"隐私求交-1": {
				Client: offlinetask.PartyInput{
					PartyID: "party-f1",
					Vtables: map[string]offlinetask.VTableInput{
						"set.table": offlinetask.VTableInput{
							VtableID:           25,
							KeyColumn:          "ID",
							SelectedVtableCols: []string{"*"},
						},
					},
				},
				Servers: []offlinetask.PartyInput{
					{
						PartyID: "party-f2",
						Vtables: map[string]offlinetask.VTableInput{
							"set.table": offlinetask.VTableInput{
								VtableID:           1,
								KeyColumn:          "ID",
								SelectedVtableCols: []string{"*"},
							},
						},
					},
				},
			},
		},
	}
	// 替换训练任务模板中的算子配置，这里只替换标签
	trainTaskParamOverwrite := &offlinetask.OperatorParams{
		Stages: map[string]offlinetask.StageParams{
			"数据分析-1": offlinetask.StageParams{
				Client: offlinetask.StagePartyParams{
					PartyID: "party-f1",
					Params: map[string]interface{}{
						"label": "income_bracket",
					},
				},
			},
			"XGB-训练": offlinetask.StageParams{
				Client: offlinetask.StagePartyParams{
					PartyID: "party-f1",
					Params: map[string]interface{}{
						"label": "income_bracket",
					},
				},
			},
		},
	}

	// 替换预测任务各参与方的输入
	predictTaskVtableInput := &offlinetask.OperatorInput{
		Stages: map[string]offlinetask.StageInput{
			"数据预处理（训练）-1": {
				Client: offlinetask.PartyInput{
					PartyID: "party-f1",
					Vtables: map[string]offlinetask.VTableInput{
						"data.table": offlinetask.VTableInput{
							VtableID:           62,
							KeyColumn:          "ID",
							SelectedVtableCols: []string{"*"},
						},
					},
				},
			},

			"数据预处理（训练）-2": {
				Client: offlinetask.PartyInput{
					PartyID: "party-f2",
					Vtables: map[string]offlinetask.VTableInput{
						"data.table": offlinetask.VTableInput{
							VtableID:           35,
							KeyColumn:          "ID",
							SelectedVtableCols: []string{"*"},
						},
					},
				},
			},
		},
	}

	//  准备训练任务==========================================================================================

	trainFileFp, err := os.OpenFile("./xgboost_train_example.json", os.O_RDONLY, os.ModePerm)
	defer trainFileFp.Close()
	if err != nil {
		// TODO: Handle error.
	}
	trainTaskJson, err := io.ReadAll(trainFileFp)
	if err != nil {
		// TODO: Handle error.
	}
	trainTask := &prototask.Task{}
	err = sudoclient.UnmarshalJSONToProto(trainTaskJson, trainTask)
	if err != nil {
		// TODO: Handle error.
	}

	sudoClientF1, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("10.0.23.130:8071"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("http://10.0.23.130:7857"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	project := offlinetask.Project{
		TusitaID:  "tusita-t1",
		ProjectID: 442450667450138891,
		PartyID:   "party-f1",
	}
	offlineTaskF1, _ := offlinetask.NewFactory(
		sudoClientF1,
		project)

	trainTask, _ = offlineTaskF1.ReplaceParty(trainTask, partyMap)
	trainTask, _ = offlineTaskF1.SetTaskInput(trainTask, &offlinetask.OperatorInput{}, trainTaskVtableInput)
	trainTask, _ = offlineTaskF1.SetTaskOperatorParams(trainTask, *trainTaskParamOverwrite)
	trainTask, _ = offlineTaskF1.SetTaskOutputDataSrc(trainTask, partyOutputDataSrc)

	// 提交训练任务
	trainTask.Name = "xg_boost_test"
	trainTask.ProjectId = offlineTaskF1.ProjectID
	createdTrainTask, err := offlineTaskF1.SudoClient.ActionOnTask(context.Background(), &task.ActionOnTaskRequest{
		Base: &prototask.ActionOnTaskRequest{
			Task: trainTask,
			Action: &prototask.ActionParam{
				Type: enums.Dag_Graph_SUBMIT,
			},
		},
	})
	if err != nil {
		// TODO: Handle error.
	}
	// 等待任务执行完成......

	// 准备预测任务========================================================================================================

	// 获取训练任务的输出模型数据，作为预测任务的输入
	modelInput, err := offlineTaskF1.GetTaskModelOutputs(context.Background(), createdTrainTask.Task.TaskId)
	if err != nil {
		// TODO: Handle error.
	}

	// 提交预测任务
	predictFileFp, err := os.OpenFile("./xgboost_predict_example.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		// TODO: Handle error.
	}
	defer predictFileFp.Close()
	predictTaskJson, err := io.ReadAll(predictFileFp)
	if err != nil {
		// TODO: Handle error.
	}

	predictTask := &prototask.Task{}

	err = sudoclient.UnmarshalJSONToProto(predictTaskJson, predictTask)
	if err != nil {
		// TODO: Handle error.
	}

	predictTask, _ = offlineTaskF1.ReplaceParty(predictTask, partyMap)
	predictTask, _ = offlineTaskF1.SetTaskOutputDataSrc(predictTask, partyOutputDataSrc)
	predictTask, _ = offlineTaskF1.SetTaskInput(predictTask, modelInput, predictTaskVtableInput)

	// 提交预测任务
	predictTask.Name = "xg_boost_predict_test"
	predictTask.ProjectId = offlineTaskF1.ProjectID
	//  必须是预测类型，后续才能创建在线服务
	predictTask.Type = enums.Job_PREDICT
	createdPredictTask, err := offlineTaskF1.SudoClient.ActionOnTask(context.Background(), &task.ActionOnTaskRequest{
		Base: &prototask.ActionOnTaskRequest{
			Task: predictTask,
			Action: &prototask.ActionParam{
				Type: enums.Dag_Graph_SUBMIT,
			},
		},
	})
	if err != nil {
		// TODO: Handle error.
	}
	// 等待任务执行完成......

	// 创建在线服务================================================================================================================
	onlinesvcF1, _ := NewFactory(sudoClientF1, project)
	sudoClientF2, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("10.0.23.130:18071"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("10.0.23.130:17857"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	// 服务方准备批准服务
	onlinesvcF2, _ := NewFactory(sudoClientF2, offlinetask.Project{
		TusitaID:  "tusita-t1",
		ProjectID: 442450667450138891,
		PartyID:   "party-f2",
	})
	// 在线服务的各方输入
	onlineTaskInput := &offlinetask.OperatorInput{
		Stages: map[string]offlinetask.StageInput{
			"数据预处理（训练）-1": {
				Client: offlinetask.PartyInput{
					PartyID: "party-f1",
					Vtables: map[string]offlinetask.VTableInput{
						"data.table": offlinetask.VTableInput{
							VtableID:  62,
							KeyColumn: "ID",
						},
					},
				},
			},

			"数据预处理（训练）-2": {
				Client: offlinetask.PartyInput{
					PartyID: "party-f2",
					Vtables: map[string]offlinetask.VTableInput{
						"data.table": offlinetask.VTableInput{
							VtableID:  35,
							KeyColumn: "ID",
						},
					},
				},
			},
		},
	}
	// 名称中含test跳过审批
	service, err := onlinesvcF1.NewService(context.TODO(), "lyy-sdk-t-est2", false, createdPredictTask.Task.TaskId, onlineTaskInput)
	if err != nil {
		// TODO: Handle error.
	}
	// 数据提供方审核
	serverService, err := onlinesvcF2.GetService(context.Background(), service.ID())
	if err != nil {
		// TODO: Handle error.
	}
	err = serverService.Approve(context.Background())
	if err != nil {
		// TODO: Handle error.
	}
	// 等待服务部署完成
	err = service.WaitForReady(context.Background())
	if err != nil {
		// TODO: Handle error.
	}
	// 使用在线服务查询
	_, err = service.Predict(context.Background(), &online_service.PredictRequest{
		Ids: []string{"223"},
	})

	if err != nil {
		// TODO: Handle error.
	}
	err = service.Delete(context.Background())
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleFactory_GetTaskVTableInputs() {
	sudoClientF1, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("10.0.23.130:8071"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("http://10.0.23.130:7857"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	f1, _ := NewFactory(
		sudoClientF1,
		offlinetask.Project{
			TusitaID:  "tusita-t1",
			ProjectID: 442450667450138891,
			PartyID:   "party-f1",
		})
	f1.GetTaskVTableInputs(context.TODO(), 18)
}
