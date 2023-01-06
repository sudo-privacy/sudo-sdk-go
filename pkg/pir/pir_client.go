package pir

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sudoprivacy.com/go/sudosdk/pkg/sudoclient"
	"sudoprivacy.com/go/sudosdk/protobuf/basic/protobuf/enums"
	"sudoprivacy.com/go/sudosdk/protobuf/virtualservice/platformpb/pir"
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
		errMsg := fmt.Sprintf("unexpected pir client service status: %s", status.String())
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
