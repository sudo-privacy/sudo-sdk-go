package sudoclient_test

import (
	"context"

	"sudoprivacy.com/go/sudosdk/pkg/sudoclient"
	"sudoprivacy.com/go/sudosdk/protobuf/virtualservice/platformpb/vtable"
)

func ExampleNewSudoClient() {
	client, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("localhost:8071"),
		sudoclient.WithHTTPEndpoint("http://localhost:7857"),
		sudoclient.WithAccount("admin", "sudosudo"))
	if err != nil {
		// TODO: Handle error.
	}
	defer client.Close()

	// 使用client调用gRPC接口
	_, err = client.ListVtables(context.TODO(), &vtable.ListVtablesRequest{})
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBasicSudoClient_CreateVtableFromLocalFile() {
	// 创建 SudoClient
	client, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("localhost:8071"),
		sudoclient.WithHTTPEndpoint("http://localhost:7857"),
		sudoclient.WithAccount("admin", "sudosudo"))
	if err != nil {
		// TODO: Handle error.
	}
	defer client.Close()
	// 从本地文件创建vtable，这里指定的是根路径下的csv表文件。
	_, err = client.CreateVtableFromLocalFile(context.TODO(), "/psi_server.csv")
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBasicSudoClient_CreateVtableFromDB() {
	client, err := sudoclient.NewSudoClient(context.TODO(),
		sudoclient.WithGrpcEndpoint("localhost:8071"),
		sudoclient.WithHTTPEndpoint("http://localhost:7857"),
		sudoclient.WithAccount("admin", "sudosudo"))
	if err != nil {
		// TODO: Handle error.
	}
	defer client.Close()
	// 指定MariaDB-server数据源，furnace_a数据库，behavior_record数据表。
	_, err = client.CreateVtableFromDB(context.TODO(), "MariaDB-server", "furnace_a", "behavior_record")
	if err != nil {
		// TODO: Handle error.
	}
}
