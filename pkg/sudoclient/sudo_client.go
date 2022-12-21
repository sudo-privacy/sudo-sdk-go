package sudoclient

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	basicplatformpb "sudosdk/protobuf/basic/protobuf/virtualservice/platformpb"
	"sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/datasource"
	basicvtable "sudosdk/protobuf/basic/protobuf/virtualservice/platformpb/vtable"
	"sudosdk/protobuf/virtualservice/platformpb"
)

type BasicSudoClient struct {
	basicplatformpb.CommonClient
	basicplatformpb.FurnaceClient
	grpcConnPool ConnPool
	httpClient   *SudoHTTPClient
}

type SudoClient struct {
	BasicSudoClient
	platformpb.CommonClient
	platformpb.FurnaceClient
	grpcConnPool ConnPool
}

func NewBasicSudoClient(ctx context.Context, opts ...DialOption) (*BasicSudoClient, error) {
	o := newDefaultdialOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	if o.tokenSource == nil {
		token, err := NewUserAccountToken(ctx, opts...)
		if err != nil {
			return nil, err
		}
		opts = append(opts, WithTokenSource(token))
	}
	grpcConnPool, err := dialPool(ctx, opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := NewHTTPClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &BasicSudoClient{
		CommonClient:  basicplatformpb.NewCommonClient(grpcConnPool),
		FurnaceClient: basicplatformpb.NewFurnaceClient(grpcConnPool),
		grpcConnPool:  grpcConnPool,
		httpClient:    httpClient,
	}, nil
}

func (client *BasicSudoClient) getIdentifyFileName(filePath string) (string, error) {
	fp, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", errors.Wrap(err, "os.OpenFile failed")
	}
	defer fp.Close()
	identityName := ""
	fileName := filepath.Base(filePath)
	hasher := sha256.New()
	_, err = io.Copy(hasher, fp)
	if err != nil {
		return "", errors.Wrap(err, "md5 hasher copy failed")
	}
	md5Str := hex.EncodeToString(hasher.Sum(nil))
	splitedFileNames := strings.Split(fileName, ".")
	for i := 0; i < len(splitedFileNames)-1; i++ {
		identityName += splitedFileNames[i]
	}
	identityName += "_" + md5Str + "." + splitedFileNames[len(splitedFileNames)-1]
	return identityName, nil
}

func (client *BasicSudoClient) CreateVtableFromLocalFile(ctx context.Context, tableFilePath string) (uint64, error) {
	identityName, err := client.getIdentifyFileName(tableFilePath)
	if err != nil {
		return 0, err
	}
	// check vtable already exist
	resp, err := client.ListVtables(ctx, &basicvtable.ListVtablesRequest{
		QueryOptions: &basicvtable.VtableQueryOptions{
			Name: identityName,
		},
	})
	if err != nil {
		return 0, errors.Wrap(err, "client.ListVtables")
	}
	for i := range resp.Data {
		if resp.Data[i] != nil {
			if resp.Data[i].Base != nil {
				if resp.Data[i].Base.Name == identityName {
					return resp.Data[i].Base.Id, nil
				}
			}
		}
	}
	// upload vtable file
	uploadedFilePath, err := client.httpClient.UploadVtableFile(ctx, tableFilePath, identityName)
	if err != nil {
		return 0, err
	}
	// send file to datasource
	updateDataSrcResp, err := client.UpdataDataSource(ctx, &datasource.UpdateDataSourceRequest{
		FilePath: uploadedFilePath,
	})
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("client.UpdataDataSource(%s)", uploadedFilePath))
	}
	splitedFileNames := strings.Split(uploadedFilePath, "/")
	// create vtable
	protoCrateVtableReq := basicvtable.CreateVtableRequest{
		Name:         identityName,
		DataShape:    "[]",
		DatasourceId: updateDataSrcResp.DatasourceId,
		Description:  identityName,
		Path:         splitedFileNames[len(splitedFileNames)-1],
	}
	createVtableResp, err := client.FurnaceClient.CreateVtable(ctx, &protoCrateVtableReq)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf(" client.FurnaceClient.CreateVtable(%+v)", protoCrateVtableReq.Name))
	}
	return createVtableResp.Data.Base.Id, nil
}

func (client *BasicSudoClient) CreateVtableFromDB(
	ctx context.Context,
	dataSrcName, dbName, tableName string,
) (uint64, error) {
	dataSrcs, err := client.GetDataSources(ctx, &datasource.GetDataSourceRequest{})
	if err != nil {
		return 0, errors.Wrap(err, "list data sources failed")
	}
	var targetDataSrcID uint64
	for _, dataSrc := range dataSrcs.Data {
		if dataSrc != nil {
			if dataSrc.Name == dataSrcName {
				targetDataSrcID = dataSrc.Id
			}
		}
	}
	if targetDataSrcID == 0 {
		return 0, errors.New("unknown dataSrcName")
	}
	vtableName := fmt.Sprintf("%s_%s_%s", dataSrcName, dbName, tableName)
	// check vtable already exist
	resp, err := client.ListVtables(ctx, &basicvtable.ListVtablesRequest{
		QueryOptions: &basicvtable.VtableQueryOptions{
			Name: vtableName,
		},
	})
	if err != nil {
		return 0, errors.Wrap(err, "client.ListVtables")
	}
	for i := range resp.Data {
		if resp.Data[i] != nil {
			if resp.Data[i].Base != nil {
				if resp.Data[i].Base.Name == tableName {
					return resp.Data[i].Base.Id, nil
				}
			}
		}
	}
	createVtableResp, err := client.CreateVtable(ctx, &basicvtable.CreateVtableRequest{
		Name:         vtableName,
		DatasourceId: targetDataSrcID,
		Description:  tableName,
		Path:         fmt.Sprintf("%s.%s", dbName, tableName),
	})
	if err != nil {
		return 0, errors.Wrap(err, "create vtable failed")
	}
	return createVtableResp.Data.Base.Id, nil
}

func (client *BasicSudoClient) Close() error {
	return client.grpcConnPool.Close()
}

func NewSudoClient(ctx context.Context, opts ...DialOption) (*SudoClient, error) {
	o := newDefaultdialOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	if o.tokenSource == nil {
		token, err := NewUserAccountToken(ctx, opts...)
		if err != nil {
			return nil, err
		}
		opts = append(opts, WithTokenSource(token))
	}
	grpcConnPool, err := dialPool(ctx, opts...)
	if err != nil {
		return nil, err
	}
	opts = append(opts, WithConnPool(grpcConnPool))
	basicSudoClient, err := NewBasicSudoClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &SudoClient{
		BasicSudoClient: *basicSudoClient,
		CommonClient:    platformpb.NewCommonClient(grpcConnPool),
		FurnaceClient:   platformpb.NewFurnaceClient(grpcConnPool),
		grpcConnPool:    grpcConnPool,
	}, nil
}

func (client *SudoClient) Close() error {
	return client.grpcConnPool.Close()
}
