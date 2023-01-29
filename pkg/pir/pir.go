package pir

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/online_service/enums"
	"github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/virtualservice/platformpb/apiusage"
	"github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/virtualservice/platformpb/pir"
	"github.com/sudo-privacy/sudo-sdk-go/pkg/sudoclient"
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
//   - 如果设置blocking，阻塞等待直到service部署完成。
//
//   - identityName 指定pir server service名称。
//
//   - svcType 设置 pir 或三要素服务。
//
//   - dataParam 配置提供服务的数据来源。
//     如果使用vtable，使用数牍隐私计算平台上已经存在的，也可以通过 [sudoclient.BasicSudoClient.CreateVtableFromLocalFile] 、
//     [sudoclient.BasicSudoClient.CreateVtableFromDB] 提前创建。
//     vtable属性可以通过
//     [github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/basic/protobuf/virtualservice/platformpb.FurnaceClient.ListVtables]
//     查询。
//
//   - keyColumns 指定pir匹配列。
//
//   - labelColumns  设置pir查询列。
//
//   - indiscernibilityDegree 设置pir不可区分度。
func (f *Factory) NewPirServer(
	ctx context.Context,
	identityName string,
	svcType enums.SVCType,
	dataParam *pir.DataModeParams,
	keyColumns []string,
	labelColumns []string,
	indiscernibilityDegree uint64,
	blocking bool,
) (*Server, error) {
	pirServer, err := f.CreatePirServer(
		ctx,
		identityName,
		svcType,
		dataParam,
		keyColumns,
		labelColumns,
		indiscernibilityDegree)
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
	svcType enums.SVCType,
	dataParam *pir.DataModeParams,
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
	for _, protoServer := range getPirServersResp.Data {
		if protoServer != nil && protoServer.Name == identityName {
			return f.protoPirServerToServer(protoServer), nil
		}
	}
	dataMode := enums.DataMode_UNKNOWN_DATA_MODE
	if dataParam.GetApiParams() != nil {
		dataMode = enums.DataMode_API
	}
	if dataParam.GetVtableParams() != nil {
		dataMode = enums.DataMode_VTABLE
	}
	if dataParam.GetMysqlParams() != nil {
		dataMode = enums.DataMode_MYSQL
	}
	if dataMode == enums.DataMode_UNKNOWN_DATA_MODE {
		return nil, errors.New("invalid dataParam")
	}

	createPirServerResp, err := f.CreatePirServerService(ctx, &pir.CreateServerServiceRequest{
		PirServer: &pir.PirServer{
			Type:                   svcType,
			DataModeParams:         dataParam,
			DataMode:               dataMode,
			IndiscernibilityDegree: indiscernibilityDegree,
			KeyColumns:             keyColumns,
			LabelColumns:           labelColumns,
			Name:                   identityName,
			Path:                   randPath(5),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err,
			fmt.Sprintf("create pir server service failed, dataParam: %+v, keyColumns: %s, LabelColumns: %s",
				dataParam, keyColumns, labelColumns))
	}
	if createPirServerResp.Data == nil {
		return nil, errors.New("invalid create pir server service response, resp.Data not found")
	}

	return f.protoPirServerModelToServer(createPirServerResp.Data), nil
}

func (f *Factory) protoPirServerModelToServer(serverModel *pir.PirServerModel) *Server {
	return &Server{
		SudoClient:   f.SudoClient,
		identityName: serverModel.Name,
		id:           serverModel.Id,
		serviceID:    serverModel.ServiceId,
	}
}

func (f *Factory) protoPirServerToServer(server *pir.PirServer) *Server {
	return &Server{
		SudoClient:   f.SudoClient,
		identityName: server.Name,
		id:           server.Id,
		serviceID:    server.ServiceId,
	}
}

// CreatePirClient 创建pir client。
// 参数配置参考 [Factory.NewPirClient] 。
func (f *Factory) CreatePirClient(
	ctx context.Context,
	serverParty, serverPath, serverServiceID, tokenStr string, svcType enums.SVCType,
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

		return f.protoPirClientToClient(data), nil
	}
	path := randPath(5)
	createPirClientServiceResp, err := f.CreatePirClientService(ctx, &pir.CreatePirClientServiceRequest{
		PirClient: &pir.PirClient{
			Name:          "pir_client_" + path,
			Type:          svcType,
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
	return f.protoPirClientModelToClient(createPirClientServiceResp.Data), nil
}

func (f *Factory) protoPirClientModelToClient(clientModel *pir.PirClientModel) *Client {
	splitPath := strings.Split(clientModel.Path, "/")
	return &Client{
		SudoClient: f.SudoClient,
		serviceID:  strconv.FormatUint(clientModel.ServiceId, 10),
		id:         clientModel.Id,
		subPath:    splitPath[len(splitPath)-1],
		svcType:    clientModel.Type,
		token:      clientModel.Token,
		name:       clientModel.Name,
	}
}

func (f *Factory) protoPirClientToClient(client *pir.PirClient) *Client {
	splitPath := strings.Split(client.Path, "/")
	return &Client{
		SudoClient: f.SudoClient,
		serviceID:  strconv.FormatUint(client.ServiceId, 10),
		id:         client.Id,
		subPath:    splitPath[len(splitPath)-1],
		svcType:    client.Type,
		token:      client.Token,
		name:       client.Name,
	}
}

// NewPirClient 创建并部署client service。
//
//   - 如果设置blocking，阻塞等待直到service部署完成。
//
//   - serverParty 设置pir服务端partyID。
//
//   - serverPath 设置服务端url地址。比如"ws://localhost:7857/ws"。
//
//   - serverServiceID 设置服务端serviceID，需要从服务端pir server读取。
//
//   - tokenStr 设置服务端提供的token，需要服务端pir server提前为客户端party创建，参考 [Server.NewClientToken] 。
//
//   - svcType 设置 pir 或三要素服务。
func (f *Factory) NewPirClient(
	ctx context.Context,
	serverParty, serverPath, serverServiceID, tokenStr string,
	svcType enums.SVCType,
	blocking bool,
) (*Client, error) {
	pirClient, err := f.CreatePirClient(ctx, serverParty, serverPath, serverServiceID, tokenStr, svcType)
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

// Manager 是对 [Factory] 的简单封装，封装管控相关接口。
type Manager struct {
	Factory
}

// NewManager 返回一个简单封装 [sudoclient.SudoClient] 的 [Manager] 。
func NewManager(factory Factory) *Manager {
	return &Manager{
		Factory: factory,
	}
}

// UpsertUsageNotifyReceiver 是对 grpc 接口的简单封装。
func (m *Manager) UpsertUsageNotifyReceiver(ctx context.Context, req *apiusage.UpsertAPIUsageReceiverRequest) error {
	resp, err := m.SudoClient.UpsertUsageNotifyReceiver(ctx, req)
	if err != nil {
		return err
	}
	if resp.Status != apiusage.UpsertAPIUsageReceiverResponse_SUCCESS {
		return status.Error(codes.Internal, resp.Msg)
	}
	return nil
}

// ListAPIUsages 是对grpc接口的简单封装。
func (m *Manager) ListAPIUsages(
	ctx context.Context,
	req *apiusage.ListAPIUsagesRequest,
) (*apiusage.ListAPIUsagesResponse, error) {
	return m.SudoClient.ListAPIUsages(ctx, req)
}

// GetServers 查询 [Server] ,根据输入的过滤参数，可能返回0个、1个 或多个 Server。
func (m *Manager) GetServers(ctx context.Context, query *pir.GetServerServicesRequest) ([]*Server, error) {
	resp, err := m.SudoClient.GetPirServerServices(ctx, query)
	if err != nil {
		return nil, err
	}
	servers := []*Server{}
	for i := range resp.Data {
		servers = append(servers, m.Factory.protoPirServerToServer(resp.Data[i]))
	}
	return servers, nil
}

// GetClients 查询 [Client] , 根据输入的过滤参数，可能返回0、1个 或多个 Client。
func (m *Manager) GetClients(ctx context.Context, query *pir.GetPirClientServicesRequst) ([]*Client, error) {
	resp, err := m.SudoClient.GetPirClientServices(ctx, query)
	if err != nil {
		return nil, err
	}
	clients := []*Client{}
	for i := range resp.Data {
		clients = append(clients, m.Factory.protoPirClientToClient(resp.Data[i]))
	}
	return clients, nil
}
