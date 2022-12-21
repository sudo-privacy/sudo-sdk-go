package sudoclient

import (
	"context"
	"sudosdk/protobuf/virtualservice/platformpb/vtable"
)

func ExampleNewClient() {
	client, err := NewSudoClient(context.TODO(), WithGrpcEndpoint("localhost:8071"), WithAccount("admin", "sudosudo"))
	if err != nil {
		// TODO: Handle error.
	}

	// do some thing with client
	_, err = client.ListVtables(context.TODO(), &vtable.ListVtablesRequest{})
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSudoClient_CreateVtableFromLocalFile() {
	// prepare SudoClient
	client, err := NewSudoClient(context.TODO(),
		WithGrpcEndpoint("localhost:8071"),
		WithHTTPEndpoint("http://localhost:7857"),
		WithAccount("admin", "sudosudo"))
	if err != nil {
		// TODO: Handle error.
	}
	// create vtable with local table file
	_, err = client.CreateVtableFromLocalFile(context.TODO(), "/psi_server.csv")
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSudoClient_CreateVtableFromDB() {
	client, err := NewSudoClient(context.TODO(),
		WithGrpcEndpoint("localhost:8071"),
		WithHTTPEndpoint("http://localhost:7857"),
		WithAccount("admin", "sudosudo"))
	if err != nil {
		// TODO: Handle error.
	}
	// create vtable with db table
	_, err = client.CreateVtableFromDB(context.TODO(), "MariaDB-server", "furnace_a", "behavior_record")
	if err != nil {
		// TODO: Handle error.
	}
}
