package onlinesvc

import (
	"context"

	"github.com/sudo-privacy/sudo-sdk-go/pkg/sudoclient"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/online_service"
)

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
	f1 := NewFactory(
		sudoClientF1,
		"tusita-t1",
		442450667450138891,
		"party-f1")

	sudoClientF2, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("10.0.23.130:18071"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("10.0.23.130:17857"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	f2 := NewFactory(sudoClientF2, "tusita-t1", 442450667450138891, "party-f2")

	input := &OperatorVtableInput{
		Stage: map[string]StageVtableInput{
			"数据预处理（训练）-1": {
				Input: map[string]PartyVtableInput{
					"party-f1": PartyVtableInput{
						Input: map[string]VTableInput{
							"data.table": VTableInput{
								VtableID:  62,
								KeyColumn: "ID",
							},
						},
					},
				},
			},
			"数据预处理（训练）-2": {
				Input: map[string]PartyVtableInput{
					"party-f2": PartyVtableInput{
						Input: map[string]VTableInput{
							"data.table": VTableInput{
								VtableID:  35,
								KeyColumn: "ID",
							},
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
	f1 := NewFactory(
		sudoClientF1,
		"tusita-t1",
		442450667450138891,
		"party-f1")
	f1.GetTaskVTableInputs(context.TODO(), 18)
}
