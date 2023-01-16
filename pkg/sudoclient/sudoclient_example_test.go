package sudoclient_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/pkg/sudoclient"
	"gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/protobuf/virtualservice/platformpb/vtable"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ExampleNewSudoClient() {
	// ==========本地无tls连接==========
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

func ExampleNewSudoClient_tls_skipVerify() {
	// ==========使用tls连接，忽略证书校验==========
	tlsCfg := tls.Config{InsecureSkipVerify: true}

	transportCred := credentials.NewTLS(&tlsCfg)
	grpcTransportOpt := grpc.WithTransportCredentials(transportCred)

	httpTransport := &http.Transport{
		TLSClientConfig: &tlsCfg,
	}

	client, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("grpc.dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("https://dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithGrpcDialOption(grpcTransportOpt),
		sudoclient.WithHTTPTransport(httpTransport),
	)
	if err != nil {
		// TODO: Handle error.
	}
	defer client.Close()
}

func ExampleNewSudoClient_tls_appendCert() {
	// ==========使用tls连接，添加自签证书到信任==========
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		// TODO: Handle error.
	}
	perm, err := ioutil.ReadFile("/home/we/dist-lyy-f1.crt")
	if err != nil {
		// TODO: Handle error.
	}
	ok := rootCAs.AppendCertsFromPEM(perm)
	if !ok {
		// TODO: Handle error.
	}
	tlsCfg := tls.Config{
		RootCAs: rootCAs,
	}
	transportCred := credentials.NewTLS(&tlsCfg)
	grpcTransportOpt := grpc.WithTransportCredentials(transportCred)
	httpTransport := &http.Transport{
		TLSClientConfig: &tlsCfg,
	}
	client, err := sudoclient.NewSudoClient(
		context.TODO(),
		sudoclient.WithGrpcEndpoint("grpc.dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithAccount("admin", "sudosudo"),
		sudoclient.WithHTTPEndpoint("https://dist-lyy-f1.dist.sudoprivacy.cn:30443"),
		sudoclient.WithGrpcDialOption(grpcTransportOpt),
		sudoclient.WithHTTPTransport(httpTransport),
	)
	if err != nil {
		// TODO: Handle error.
	}
	defer client.Close()
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
