// Package sudoclient 封装sudo-api生成的gRPC client。
//
// [SudoClient] 自动处理token的刷新，支持连接池配置，简化gRPC接口使用。
//
// # Client 使用
//
// 下面是创建SudoClient并使用sudo-api生成的gRPC client接口的example。
//
//	func ExampleNewSudoClient() {
//		client, err := NewSudoClient(context.TODO(),
//			WithGrpcEndpoint("localhost:8071"),
//			WithHTTPEndpoint("http://localhost:7857"),
//			WithAccount("admin", "sudosudo"))
//		if err != nil {
//			// TODO: Handle error.
//		}
//		defer client.Close()
//
//		// 使用 client
//		_, err = client.ListVtables(context.TODO(), &vtable.ListVtablesRequest{})
//		if err != nil {
//			// TODO: Handle error.
//		}
//	}
//
// # Context 使用
//
// 传入NewSudoClient的context会用于请求token和创建底层连接。
// client的每个方法调用可能会触发token刷新。刷新token也是gRPC接口调用，用于token刷新的grpc接口调用会使用方法传入的context。
package sudoclient
