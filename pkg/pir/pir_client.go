package pir

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sudo-privacy/sudo-sdk-go/pkg/sudoclient"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/online_service"
	onlinesvcenums "github.com/sudo-privacy/sudo-sdk-go/protobuf/online_service/enums"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/pir"
)

var clientStableStatus = []enums.PirService_Status{
	enums.PirService_UNKNOWN_STATUS,
	enums.PirService_READY_TO_BOOTING,
	enums.PirService_ACTIVE,
	enums.PirService_USING,
	enums.PirService_FAILED,
	enums.PirService_TERMINATED,
}

// Client 封装了关于pir client service的gRPC接口，可用于client service管理和pir查询。
type Client struct {
	*sudoclient.SudoClient
	serviceID string
	token     string
	id        uint64
	subPath   string
	svcType   onlinesvcenums.SVCType
	name      string
}

// Deploy 部署pir service。
// 如果blocking为true，阻塞等待直到服务部署完成。
func (c *Client) Deploy(ctx context.Context, blocking bool) error {
	_, err := c.PirClientServiceTakeAction(ctx, &pir.ClientServiceTakeActionRequest{
		ServiceIdStr: c.serviceID,
		Token:        c.token,
		Action:       enums.PirService_DEPLOY.String(),
	})
	if err != nil {
		return errors.Wrap(err, "deploy pir client service failed")
	}
	if blocking {
		err = c.WaitForReady(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) getStatus(ctx context.Context) (enums.PirService_Status, error) {
	resp, err := c.GetPirClientServices(ctx, &pir.GetPirClientServicesRequst{
		Query: &pir.ClientQueryOption{
			Id: c.id,
		},
	})
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("get pir client service failed, serviceID:%d", c.id))
	}

	if len(resp.Data) == 0 {
		notFoundErr := status.Error(codes.NotFound, "pir client service not found")
		return 0, errors.Wrap(notFoundErr, "no data in get pir client service response")
	}
	if len(resp.Data) > 1 {
		return 0, errors.New("multi pir client service found with same service id")
	}
	if resp.Data[0] == nil {
		return 0, errors.New("internal error,response Data not exist")
	}
	return resp.Data[0].Status, nil
}

func (c *Client) getMessage(ctx context.Context) (string, error) {
	resp, err := c.GetPirClientServices(ctx, &pir.GetPirClientServicesRequst{
		Query: &pir.ClientQueryOption{
			Id: c.id,
		},
	})
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("get pir client service failed, serviceID:%d", c.id))
	}

	if len(resp.Data) == 0 {
		notFoundErr := status.Error(codes.NotFound, "pir client service not found")
		return "", errors.Wrap(notFoundErr, "no data in get pir client service response")
	}
	if len(resp.Data) > 1 {
		return "", errors.New("multi pir client service found with same service id")
	}
	if resp.Data[0] == nil {
		return "", errors.New("internal error,response Data not exist")
	}
	return resp.Data[0].Message, nil
}

// IsActive 检查service 是否就绪可用。
func (c *Client) IsActive(ctx context.Context) (bool, error) {
	status, err := c.getStatus(ctx)
	if err != nil {
		return false, err
	}
	if status == enums.PirService_ACTIVE || status == enums.PirService_USING {
		return true, nil
	}
	return false, nil
}

func (c *Client) waitStatus(
	ctx context.Context,
	targetStatus []enums.PirService_Status,
	interval time.Duration,
) (enums.PirService_Status, error) {
	status, err := c.getStatus(ctx)
	if err != nil {
		return 0, err
	}
	for i := range targetStatus {
		if status == targetStatus[i] {
			return targetStatus[i], nil
		}
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			status, err := c.getStatus(ctx)
			if err != nil {
				return 0, err
			}
			for i := range targetStatus {
				if status == targetStatus[i] {
					return targetStatus[i], nil
				}
			}
		case <-ctx.Done():
			return 0, errors.New("wait status cxt Done")
		}
	}
}

// WaitForReady 阻塞等待直到 service 部署完成，如果检查到部署失败会直接返回错误。
func (c *Client) WaitForReady(ctx context.Context) error {
	status, err := c.waitStatus(ctx, clientStableStatus, waitInterval)
	if err != nil {
		return err
	}
	if status != enums.PirService_ACTIVE {
		message, err := c.getMessage(ctx)
		if err != nil {
			return err
		}
		errMsg := fmt.Sprintf("unexpected pir client service status: %s,Message: %s", status.String(), message)
		return errors.New(errMsg)
	}
	return nil
}

// Delete 终结并删除service。
func (c *Client) Delete(ctx context.Context) error {
	currentStatus, err := c.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currentStatus != enums.PirService_TERMINATED && currentStatus != enums.PirService_FAILED {
		err = c.TerminateService(ctx)
		if err != nil {
			return errors.Wrap(err, "terminate pir client service failed")
		}
	}

	err = c.deleteService(ctx)
	if err != nil {
		return err
	}
	return nil
}

// TerminateService 仅终结service。service未删除，之后还可以重新部署。
func (c *Client) TerminateService(ctx context.Context) error {
	_, err := c.PirClientServiceTakeAction(ctx, &pir.ClientServiceTakeActionRequest{
		ServiceIdStr: c.serviceID,
		Token:        c.token,
		Action:       enums.PirService_TERMINATE.String(),
	})
	if err != nil {
		return errors.Wrap(err, "terminate pir client service failed")
	}
	return nil
}

func (c *Client) deleteService(ctx context.Context) error {
	currentStatus, err := c.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currentStatus != enums.PirService_TERMINATED && currentStatus != enums.PirService_FAILED {
		return errors.New("can't delete pir client service,service not terminated or failed")
	}
	_, err = c.DeletePirClientService(ctx, &pir.DeleteClientServiceRequest{
		ServiceIdStr: c.serviceID,
		Token:        c.token,
	})
	if err != nil {
		return errors.Wrap(err, "delete pir client service failed")
	}
	return nil
}

// Pir 查询。
func (c *Client) Pir(ctx context.Context, queries []*pir.PirRequest_KeyColumn) (*pir.PirResponse, error) {
	if c.svcType != *onlinesvcenums.SVCType_PIR.Enum() {
		return nil, errors.New("not pir client service")
	}
	serviceIDUint64, err := strconv.ParseUint(c.serviceID, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("invalid serverServiceID: %s", c.serviceID)
		return nil, errors.New(errMsg)
	}
	pirResp, err := c.PirRequest(ctx, &pir.PirRequest{
		ServiceId: serviceIDUint64,
		Token:     c.token,
		Keys:      queries,
		SubPath:   c.subPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Pir request failed")
	}
	return pirResp, nil
}

// Factor3SVerify 三要素查询，grpc接口的简单封装，query参数不用设置subPath字段。
func (c *Client) Factor3SVerify(
	ctx context.Context,
	query *online_service.Factor3SVerifyRequest,
) (*online_service.Factor3SVerifyResponse, error) {
	if c.svcType != *onlinesvcenums.SVCType_FACTOR3s.Enum() {
		return nil, errors.New("not FACTOR3s client service")
	}
	query.SubPath = c.subPath
	return c.SudoClient.Factor3SVerifyQuery(ctx, query)
}

// Status 查询并返回状态。
func (c *Client) Status(ctx context.Context) (enums.PirService_Status, error) {
	return c.getStatus(ctx)
}

// ID 返回id属性。
func (c *Client) ID() uint64 {
	return c.id
}

// Name 返回name属性
func (c *Client) Name() string {
	return c.name
}

// ServiceID 返回serviceID属性
func (c *Client) ServiceID() string {
	return c.serviceID
}

// Token 返回token属性。
func (c *Client) Token() string {
	return c.token
}
