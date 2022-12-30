package pir_test

import (
	"context"
	"fmt"

	"sudoprivacy.com/go/sudosdk/pkg/pir"
	"sudoprivacy.com/go/sudosdk/pkg/sudoclient"
	"sudoprivacy.com/go/sudosdk/protobuf/online_service"
)

func ExampleClient_Pir() {
	//=================server side====================

	// prepare server side sudoClient
	sudoClientF1, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("localhost:8071"),
		sudoclient.WithHTTPEndpoint("http://localhost:7857"),
		sudoclient.WithAccount("admin", "sudosudo"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	factoryF1 := pir.NewPirFactory(sudoClientF1)
	// New pir pirServerF1(blocking,idempotent by identityName)
	pirServerF1, err := factoryF1.NewPirServer(context.TODO(),
		"IDENTITY_NAME_HERE",
		7,
		[]string{"ID"},
		[]string{"age"},
		10,
		true,
	)
	// prepare client token
	token, err := pirServerF1.NewClientToken(context.TODO(), "party-f2")
	serviceID := pirServerF1.GetServiceID()

	//=================client side====================

	// prepare client side sudoClient
	sudoClientF2, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("localhost:18071"),
		sudoclient.WithHTTPEndpoint("http://localhost:17857"),
		sudoclient.WithAccount("admin", "sudosudo"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	factoryF2 := pir.NewPirFactory(sudoClientF2)
	// new pirClientF2(blocking,idempotent by token)
	pirClientF2, err := factoryF2.NewPirClient(context.TODO(),
		"party-f1",
		"ws://localhost:7857/ws", // 服务方nginx地址，https入口需配置为wss
		fmt.Sprintf("%d", serviceID),
		token,
		true,
	)
	if err != nil {
		// TODO: Handle error.
	}

	query := []*online_service.KeyColumn{}
	query = append(query, &online_service.KeyColumn{
		Values: []string{"13380"},
	})
	query = append(query, &online_service.KeyColumn{
		Values: []string{"17816"},
	})
	// do pir request
	_, err = pirClientF2.Pir(context.TODO(), query)
	if err != nil {
		// TODO: Handle error.
	}

	// ============release all resoure===============
	err = pirClientF2.Delete(context.TODO())
	err = pirServerF1.Delete(context.TODO())
	factoryF1.Close()
	factoryF2.Close()
}
