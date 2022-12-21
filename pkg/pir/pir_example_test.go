package pir

import (
	"context"
	"fmt"
	"sudosdk/pkg/sudoclient"
	"sudosdk/protobuf/online_service"
)

func ExampleClient_Pir() {
	//=================server side====================

	// prepare server side sudoClient
	client1, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("localhost:8071"),
		sudoclient.WithHTTPEndpoint("http://localhost:7857"),
		sudoclient.WithAccount("admin", "sudosudo"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	factory1 := NewPirFactory(client1)
	// New pir server(blocking,idempotent by identityName)
	server, err := factory1.NewPirServer(context.TODO(),
		"IDENTITY_NAME_HERE",
		7,
		[]string{"ID"},
		[]string{"age"},
		10,
		true,
	)
	// prepare client token
	token, err := server.NewClientToken(context.TODO(), "party-f2")
	serviceID := server.GetServiceID()

	//=================client side====================

	// prepare client side sudoClient
	client2, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("localhost:18071"),
		sudoclient.WithHTTPEndpoint("http://localhost:17857"),
		sudoclient.WithAccount("admin", "sudosudo"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	factory2 := NewPirFactory(client2)
	// new client(blocking,idempotent by token)
	client, err := factory2.NewPirClient(context.TODO(),
		"party-f1",
		"ws://localhost:7857/ws", // 服务方nginx地址
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
	_, err = client.Pir(context.TODO(), query)
	if err != nil {
		// TODO: Handle error.
	}

	// ============release all resoure===============
	err = client.Delete(context.TODO())
	err = server.Delete(context.TODO())

}
