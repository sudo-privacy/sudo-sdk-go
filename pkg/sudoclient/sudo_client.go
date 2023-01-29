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

	"github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/basic/protobuf/virtualservice/platformpb/datasource"
	basiccommon "github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/basic/protobuf/virtualservice/platformpb/service/common"
	basicfurnace "github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/basic/protobuf/virtualservice/platformpb/service/furnace"
	basicvtable "github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/basic/protobuf/virtualservice/platformpb/vtable"
	common "github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/virtualservice/platformpb/service/common"
	furnace "github.com/sudo-privacy/sudo-sdk-go/pkg/protobuf/virtualservice/platformpb/service/furnace"
)

// BasicSudoClient 封装基础版sudo-api生成的 gRPC client。
type BasicSudoClient struct {
	basiccommon.CommonClient
	basicfurnace.FurnaceClient
	grpcConnPool ConnPool
	httpClient   *SudoHTTPClient
}

// SudoClient 封装标准版sudo-api生成的 gRPC client。
type SudoClient struct {
	BasicSudoClient
	common.CommonClient
	furnace.FurnaceClient
	grpcConnPool ConnPool
}

// NewBasicSudoClient 创建基础版SudoClient，其中仅包含基础版sudo-api提供的gRPC client。
// client使用后需要关闭，从而释放其中的连接资源。
func NewBasicSudoClient(ctx context.Context, opts ...DialOption) (*BasicSudoClient, error) {
	o := newDefaultDialOptions()
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
		CommonClient:  basiccommon.NewCommonClient(grpcConnPool),
		FurnaceClient: basicfurnace.NewFurnaceClient(grpcConnPool),
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
	splitFileNames := strings.Split(fileName, ".")
	for i := 0; i < len(splitFileNames)-1; i++ {
		identityName += splitFileNames[i]
	}
	identityName += "_" + md5Str + "." + splitFileNames[len(splitFileNames)-1]
	return identityName, nil
}

// CreateVtableFromLocalFile 从本地文件创建vtable，返回vtableID。
// 重复上传相同的文件会直接返回之前的vtableID。
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
	updateDataSrcResp, err := client.FurnaceClient.SendFileToDataSource(ctx, &datasource.SendFileToDataSourceRequest{
		FilePath: uploadedFilePath,
	})
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("client.UpdataDataSource(%s)", uploadedFilePath))
	}
	splitFileNames := strings.Split(uploadedFilePath, "/")
	// create vtable
	protoCrateVtableReq := basicvtable.CreateVtableRequest{
		Name:         identityName,
		DataShape:    "[]",
		DatasourceId: updateDataSrcResp.DatasourceId,
		Description:  identityName,
		Path:         splitFileNames[len(splitFileNames)-1],
	}
	createVtableResp, err := client.FurnaceClient.CreateVtable(ctx, &protoCrateVtableReq)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf(" client.FurnaceClient.CreateVtable(%+v)", protoCrateVtableReq.Name))
	}
	return createVtableResp.Data.Base.Id, nil
}

// CreateVtableFromDB 指定数据表创建vtable，返回vtableID。
// dataSource需要在数牍隐私计算平台提前配置，dataSource名称会被校验是否存在。
// 重复创建相同的数据表会直接返回之前的vtableID。
func (client *BasicSudoClient) CreateVtableFromDB(
	ctx context.Context,
	dataSrcName, dbName, tableName string,
) (uint64, error) {
	dataSrcs, err := client.GetDataSources(ctx, &datasource.GetDataSourcesRequest{})
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

// Close 关闭其中的连接资源。
func (client *BasicSudoClient) Close() error {
	return client.grpcConnPool.Close()
}

// NewSudoClient 创建标准版sudo-api gRPC client。SudoClient继承了BasicSudoClient的全部功能。
// client使用后需要关闭，从而释放其中的连接资源。
func NewSudoClient(ctx context.Context, opts ...DialOption) (*SudoClient, error) {
	o := newDefaultDialOptions()
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
		CommonClient:    common.NewCommonClient(grpcConnPool),
		FurnaceClient:   furnace.NewFurnaceClient(grpcConnPool),
		grpcConnPool:    grpcConnPool,
	}, nil
}

// Close 关闭其中的连接资源。
func (client *SudoClient) Close() error {
	return client.grpcConnPool.Close()
}
