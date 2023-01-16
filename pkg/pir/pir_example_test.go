package pir_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/pkg/pir"
	"gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/pkg/sudoclient"
	"gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/protobuf/online_service"
	"gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/protobuf/online_service/enums"
	protopir "gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/protobuf/virtualservice/platformpb/pir"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
		enums.SVCType_PIR,
		&protopir.DataModeParams{
			Params: &protopir.DataModeParams_VtableParams{
				VtableParams: &protopir.VtableModeParams{
					VtableId: 7,
				},
			},
		},
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
		enums.SVCType_PIR,
		true,
	)
	if err != nil {
		// TODO: Handle error.
	}

	query := []*protopir.PirRequest_KeyColumn{}
	query = append(query, &protopir.PirRequest_KeyColumn{
		Values: []string{"13380"},
	})
	query = append(query, &protopir.PirRequest_KeyColumn{
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

func ExampleClient_Factor3SVerify() {
	tlsCfg := tls.Config{InsecureSkipVerify: true}

	transportCred := credentials.NewTLS(&tlsCfg)
	grpcTransportOpt := grpc.WithTransportCredentials(transportCred)

	httpTransport := &http.Transport{
		TLSClientConfig: &tlsCfg,
	}

	//=================server side====================

	// prepare server side sudoClient
	client1, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("grpc.dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithHTTPEndpoint("https://dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithGrpcDialOption(grpcTransportOpt),
		sudoclient.WithHTTPTransport(httpTransport),
	)
	if err != nil {
		// TODO: Handle error.
	}
	factoryF1 := pir.NewPirFactory(client1)
	// New pir server(blocking,idempotent by identityName)
	server, err := factoryF1.NewPirServer(context.TODO(),
		"lyy-server-test11",
		enums.SVCType_FACTOR3s,
		&protopir.DataModeParams{
			Params: &protopir.DataModeParams_ApiParams{
				ApiParams: &protopir.APIModeParams{
					Url:  "http://10.0.1.40:9123",
					Path: "/api/mock/data",
					AuthParams: &protopir.AuthParams{
						AuthMethod: enums.AuthMethod_NO_AUTH,
					},
				},
			},
		},
		[]string{"phone_md5"},
		[]string{"id_md5", "name_md5"},
		10,
		true,
	)
	if err != nil {
		// TODO: Handle error.
	}
	// prepare client token
	token, err := server.NewClientToken(context.TODO(), "dist-lyy-f2")
	if err != nil {
		// TODO: Handle error.
	}
	serviceID := server.GetServiceID()

	//=================client side====================

	// prepare client side sudoClient
	clientF2, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("grpc.dist-lyy-f2.dist.sudoprivacy.cn:30443"),
		sudoclient.WithHTTPEndpoint("https://dist-lyy-f2.dist.sudoprivacy.cn:30443"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithGrpcDialOption(grpcTransportOpt),
		sudoclient.WithHTTPTransport(httpTransport),
	)
	if err != nil {
		// TODO: Handle error.
	}
	factoryF2 := pir.NewPirFactory(clientF2)
	// new client(blocking,idempotent by token)
	client, err := factoryF2.NewPirClient(context.TODO(),
		"dist-lyy-f1",
		"wss://dist-lyy-f1.dist.sudoprivacy.cn:30443/ws", // 如果是 http 则配置为ws
		fmt.Sprintf("%d", serviceID),
		token,
		enums.SVCType_FACTOR3s,
		true,
	)
	if err != nil {
		// TODO: Handle error.
	}
	// 三要素查询
	query := &online_service.Factor3SVerifyRequest{
		PhoneMd5: "63b78922bef3d161d2e617080e233011",
		IdMd5:    "4c28b1a6e8bb7cdc6d6050c61a748802",
		NameMd5:  "3a30b24a583acdb203ab1092c793cdc2",
	}
	_, err = client.Factor3SVerify(context.TODO(), query)
	if err != nil {
		// TODO: Handle error.
	}

	// ============release all resoure===============
	err = client.Delete(context.TODO())
	err = server.Delete(context.TODO())
	factoryF1.Close()
	factoryF2.Close()
}

func ExampleServer_UpdateServerData() {
	tlsCfg := tls.Config{InsecureSkipVerify: true}

	transportCred := credentials.NewTLS(&tlsCfg)
	grpcTransportOpt := grpc.WithTransportCredentials(transportCred)

	httpTransport := &http.Transport{
		TLSClientConfig: &tlsCfg,
	}

	client1, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("grpc.dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithHTTPEndpoint("https://dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithGrpcDialOption(grpcTransportOpt),
		sudoclient.WithHTTPTransport(httpTransport),
	)
	if err != nil {
		// TODO: Handle error.
	}
	factory1 := pir.NewPirFactory(client1)
	manager := pir.NewManager(*factory1)
	servers, err := manager.GetServers(context.TODO(), &protopir.GetServerServicesRequest{
		Query: &protopir.ServerQueryOption{
			Id: 14,
		},
	})
	if err != nil {
		// TODO: Handle error.
	}
	err = servers[0].UpdateServerData(context.TODO(), &protopir.DataModeParams{
		Params: &protopir.DataModeParams_VtableParams{
			VtableParams: &protopir.VtableModeParams{
				// new Vtable
				VtableId: 3,
			},
		},
	})
	if err != nil {
		// TODO: Handle error.
	}
}
