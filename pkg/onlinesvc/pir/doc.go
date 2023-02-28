// Package pir 简化pir服务的部署和使用，基于 [sudoclient.SudoClient] 封装的gRPC接口实现。
//
// [Server] 、[Client] 分别封装了pir server，pir client相关的gRPC接口调用。
//
// # 部署 Pir Server Example
//
// 下面是创建并部署pir server service的example。
//
//	clientF1, err := sudoclient.NewSudoClient(context.TODO(),
//		sudoclient.WithGrpcEndpoint("localhost:8071"),
//		sudoclient.WithHTTPEndpoint("http://localhost:7857"),
//		sudoclient.WithAccount("admin", "sudosudo"),
//	)
//	if err != nil {
//		// TODO: Handle error.
//	}
//	factoryF1 := NewPirFactory(clientF1)
//	// 创建并等待pir server service部署完成
//	pirServerF1, err := factoryF1.NewPirServer(context.TODO(),
//		"IDENTITY_NAME_HERE",
//		enums.SVCType_PIR,
//		&protopir.DataModeParams{
//			Params: &protopir.DataModeParams_VtableParams{
//				VtableParams: &protopir.VtableModeParams{
//					VtableId: 7,
//				},
//			},
//		},
//		[]string{"ID"},
//		[]string{"age"},
//		10,
//		true,
//	)
//
package pir

import (
	// Register fo doc
	_ "github.com/sudo-privacy/sudo-sdk-go/pkg/sudoclient"
)
