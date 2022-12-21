package pir

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"sudosdk/pkg/sudoclient"
	"sudosdk/protobuf/virtualservice/platformpb/pir"
)

type Factory struct {
	*sudoclient.SudoClient
}

func NewPirFactory(client *sudoclient.SudoClient) *Factory {
	return &Factory{
		SudoClient: client,
	}
}

func (f *Factory) NewPirServer(
	ctx context.Context,
	identityName string,
	vtableID uint64,
	keyColumes []string,
	labelColumes []string,
	indiscernibilityDegree uint64,
	blocking bool,
) (*Server, error) {
	pirServer, err := f.CreatePirServer(ctx, identityName, vtableID, keyColumes, labelColumes, indiscernibilityDegree)
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

func (f *Factory) CreatePirServer(
	ctx context.Context,
	identityName string,
	vtableID uint64,
	keyColumes []string,
	labelColumes []string,
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
			KeyColumns:             keyColumes,
			LabelColumns:           labelColumes,
			Name:                   identityName,
			VtableId:               vtableID,
			Path:                   randPath(5),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err,
			fmt.Sprintf("create pir server service failed, vtableID: %d, keyColumes: %s, lable_columes: %s",
				vtableID, keyColumes, labelColumes))
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
		splitedPath := strings.Split(data.Path, "/")
		return &Client{
			SudoClient: f.SudoClient,
			id:         data.Id,
			token:      tokenStr,
			serviceID:  serverServiceID,
			subPath:    splitedPath[len(splitedPath)-1],
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
