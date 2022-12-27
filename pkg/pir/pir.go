package pir

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"sudoprivacy.com/go/sudosdk/pkg/sudoclient"
	"sudoprivacy.com/go/sudosdk/protobuf/virtualservice/platformpb/pir"
)

// Factory 是对 [sudoclient.SudoClient] 的简单封装，提供创建pir Server/Client方法。
type Factory struct {
	*sudoclient.SudoClient
}

// NewPirFactory 返回一个简单封装 [sudoclient.SudoClient] 的 [Factory] 。
func NewPirFactory(client *sudoclient.SudoClient) *Factory {
	return &Factory{
		SudoClient: client,
	}
}

// NewPirServer 创建并部署 pir server。
//
//	- 如果设置blocking，阻塞等待直到service部署完成。
//
//	- identityName 指定pir server service名称。
//
//	- vtableID 指定提供数据服务的vtable。
//	  vtable可以使用数牍隐私计算平台上已经存在的，也可以通过 [sudoclient.BasicSudoClient.CreateVtableFromLocalFile] 、
//	   [sudoclient.BasicSudoClient.CreateVtableFromDB] 提前创建。
//	  vtable属性可以通过
//	   [sudoprivacy.com/go/sudosdk/protobuf/basic/protobuf/virtualservice/platformpb.FurnaceClient.ListVtables]  查询。
//
//	- keyColumns 指定pir匹配列。
//
//	- labelColumns  设置pir查询列。
//
//	- indiscernibilityDegree 设置pir不可区分度。
func (f *Factory) NewPirServer(
	ctx context.Context,
	identityName string,
	vtableID uint64,
	keyColumns []string,
	labelColumns []string,
	indiscernibilityDegree uint64,
	blocking bool,
) (*Server, error) {
	pirServer, err := f.CreatePirServer(ctx, identityName, vtableID, keyColumns, labelColumns, indiscernibilityDegree)
	if err != nil {
		return nil, err
	}
	active, err := pirServer.IsActive(ctx)
	if err != nil {
		return nil, err
	}
	if !active {
		err = pirServer.Deploy(ctx, blocking)
		if err != nil {
			return nil, err
		}
	}
	return pirServer, nil
}

// CreatePirServer 创建pir server。
// 参数配置参考 [Factory.NewPirServer] 。
func (f *Factory) CreatePirServer(
	ctx context.Context,
	identityName string,
	vtableID uint64,
	keyColumns []string,
	labelColumns []string,
	indiscernibilityDegree uint64,
) (*Server, error) {
	getPirServersResp, err := f.GetPirServerServices(ctx, &pir.GetServerServicesRequest{
		Query: &pir.ServerQueryOption{},
	})
	if err != nil {
		return nil, errors.Wrap(err, "list pirServer failed")
	}
	for _, server := range getPirServersResp.Data {
		if server != nil && server.Name == identityName {
			return &Server{
				SudoClient: f.SudoClient,
				id:         server.Id,
				serviceID:  server.ServiceId,
			}, nil
		}
	}
	createPirServerResp, err := f.CreatePirServerService(ctx, &pir.CreateServerServiceRequest{
		PirServer: &pir.PirServer{
			IndiscernibilityDegree: indiscernibilityDegree,
			KeyColumns:             keyColumns,
			LabelColumns:           labelColumns,
			Name:                   identityName,
			VtableId:               vtableID,
			Path:                   randPath(5),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err,
			fmt.Sprintf("create pir server service failed, vtableID: %d, keyColumns: %s, LabelColumns: %s",
				vtableID, keyColumns, labelColumns))
	}
	if createPirServerResp.Data == nil {
		return nil, errors.New("invalid create pir server service response, resp.Data not found")
	}
	return &Server{
		SudoClient:   f.SudoClient,
		identityName: identityName,
		id:           createPirServerResp.Data.Id,
		serviceID:    createPirServerResp.Data.ServiceId,
	}, nil
}

// CreatePirClient 创建pir client。
// 参数配置参考 [Factory.NewPirClient] 。
func (f *Factory) CreatePirClient(
	ctx context.Context,
	serverParty, serverPath, serverServiceID, tokenStr string,
) (*Client, error) {
	serviceIDUint64, err := strconv.ParseUint(serverServiceID, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("invalid serverServiceID: %s", serverServiceID))
	}
	getPirClientServiceResp, err := f.GetPirClientServices(ctx, &pir.GetPirClientServicesRequst{
		Query: &pir.ClientQueryOption{
			ServiceId: serviceIDUint64,
			Token:     tokenStr,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("get pir client services with token %s failed", tokenStr))
	}
	if len(getPirClientServiceResp.Data) > 1 {
		errMsg := fmt.Sprintf("found multi pir client service with token: %s", tokenStr)
		return nil, errors.New(errMsg)
	}
	if len(getPirClientServiceResp.Data) == 1 {
		data := getPirClientServiceResp.Data[0]
		if data == nil {
			return nil, errors.New("internal error, get pir client services return nil")
		}
		splitPath := strings.Split(data.Path, "/")
		return &Client{
			SudoClient: f.SudoClient,
			id:         data.Id,
			token:      tokenStr,
			serviceID:  serverServiceID,
			subPath:    splitPath[len(splitPath)-1],
		}, nil
	}
	path := randPath(5)
	createPirClientServiceResp, err := f.CreatePirClientService(ctx, &pir.CreatePirClientServiceRequest{
		PirClient: &pir.PirClient{
			Name:          "pir_client_" + path,
			Path:          path,
			ServerPartyId: serverParty,
			ServerPath:    serverPath,
			Token:         tokenStr,
		},
		ServiceIdStr: serverServiceID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create pir client service failed")
	}
	if createPirClientServiceResp.Data == nil {
		return nil, errors.New("internal error, response nil data")
	}
	return &Client{
		SudoClient: f.SudoClient,
		serviceID:  serverServiceID,
		id:         createPirClientServiceResp.Data.Id,
		subPath:    path,
		token:      tokenStr,
	}, nil
}

// NewPirClient 创建并部署client service。
//
//	- 如果设置blocking，阻塞等待直到service部署完成。
//
//	- serverParty 设置pir服务端partyID。
//
//	- serverPath 设置服务端url地址。比如"ws://localhost:7857/ws"。
//
//	- serverServiceID 设置服务端serviceID，需要从服务端pir server读取。
//
//	- tokenStr 设置服务端提供的token，需要服务端pir server提前为客户端party创建，参考 [Server.NewClientToken] 。
func (f *Factory) NewPirClient(
	ctx context.Context,
	serverParty, serverPath, serverServiceID, tokenStr string,
	blocking bool,
) (*Client, error) {
	pirClient, err := f.CreatePirClient(ctx, serverParty, serverPath, serverServiceID, tokenStr)
	if err != nil {
		return nil, err
	}
	active, err := pirClient.IsActive(ctx)
	if err != nil {
		return nil, err
	}
	if !active {
		err = pirClient.Deploy(ctx, blocking)
		if err != nil {
			return nil, err
		}
	}
	return pirClient, nil
}
