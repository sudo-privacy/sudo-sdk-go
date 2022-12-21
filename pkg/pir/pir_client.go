package pir

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sudosdk/pkg/sudoclient"
	"sudosdk/protobuf/basic/protobuf/enums"
	"sudosdk/protobuf/online_service"
	"sudosdk/protobuf/virtualservice/platformpb/pir"
)

var clientStableStatus = []enums.PirService_Status{
	enums.PirService_UNKNOWN_STATUS,
	enums.PirService_READY_TO_BOOTING,
	enums.PirService_ACTIVE,
	enums.PirService_USING,
	enums.PirService_FAILED,
	enums.PirService_TERMINATED,
}

type Client struct {
	*sudoclient.SudoClient
	serviceID string
	token     string
	id        uint64
	subPath   string
}

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

func (c *Client) Delete(ctx context.Context) error {
	currStatus, err := c.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currStatus != enums.PirService_TERMINATED && currStatus != enums.PirService_FAILED {
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

// TerminateService 返回后状态立刻变成Terminated,无需等待
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
	currStatus, err := c.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currStatus != enums.PirService_TERMINATED && currStatus != enums.PirService_FAILED {
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

func (c *Client) Pir(ctx context.Context, queries []*online_service.KeyColumn) (*pir.PirResponse, error) {
	serviceIDUint64, err := strconv.ParseUint(c.serviceID, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("invalid serverServiceID: %s", c.serviceID)
		return nil, errors.New(errMsg)
	}
	pirResp, err := c.PirRequest(ctx, &pir.PirRequest{
		Base: &online_service.PirRequest{
			ServiceId: serviceIDUint64,
			Token:     c.token,
			Keys:      queries,
		},
		SubPath: c.subPath,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Pir request failed")
	}
	return pirResp, nil
}
