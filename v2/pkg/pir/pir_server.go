package pir

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sudo-privacy/sudo-sdk-go/v2/pkg/sudoclient"
	"github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/enums"
	"github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/pir"
	"github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/token"
)

var waitInterval = time.Second * 5

var serverStableStatus = []enums.PirService_Status{
	enums.PirService_UNKNOWN_STATUS,
	enums.PirService_READY_TO_BOOTING,
	enums.PirService_ACTIVE,
	enums.PirService_USING,
	enums.PirService_FAILED,
	enums.PirService_TERMINATED,
}

// Server 封装了关于pir server service的gRPC接口。
// 用于service的部署、终结，和token创建、停用等操作。
type Server struct {
	*sudoclient.SudoClient
	identityName string
	id           uint64
	serviceID    uint64
}

// Deploy 部署service.
// 如果设置blocking,阻塞等待直到服务就绪。
func (s *Server) Deploy(ctx context.Context, blocking bool) error {
	_, err := s.PirServerServiceTakeAction(ctx, &pir.ServerServiceTakeActionRequest{
		ServiceIdStr: fmt.Sprintf("%d", s.serviceID),
		Action:       enums.PirService_DEPLOY.String(),
	})
	if err != nil {
		return errors.Wrap(err, "deploy pir server failed")
	}
	if blocking {
		err = s.WaitForReady(ctx, waitInterval)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) getStatus(ctx context.Context) (enums.PirService_Status, error) {
	resp, err := s.GetPirServerServices(ctx, &pir.GetServerServicesRequest{
		Query: &pir.ServerQueryOption{
			ServiceId: s.serviceID,
		},
	})
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("get pir server service failed,serviceID:%d", s.serviceID))
	}

	if len(resp.Data) == 0 {
		notFoundErr := status.Error(codes.NotFound, "pir server service not found")
		return 0, errors.Wrap(notFoundErr, "no data in get pir server service response")
	}
	if len(resp.Data) > 1 {
		return 0, errors.New("multi pir server service found with same service id")
	}
	if resp.Data[0] == nil {
		return 0, errors.New("internal error,response Data not exist")
	}
	return resp.Data[0].Status, nil
}

func (s *Server) getMessage(ctx context.Context) (string, error) {
	resp, err := s.GetPirServerServices(ctx, &pir.GetServerServicesRequest{
		Query: &pir.ServerQueryOption{
			ServiceId: s.serviceID,
		},
	})
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("get pir server service failed,serviceID:%d", s.serviceID))
	}

	if len(resp.Data) == 0 {
		notFoundErr := status.Error(codes.NotFound, "pir server service not found")
		return "", errors.Wrap(notFoundErr, "no data in get pir server service response")
	}
	if len(resp.Data) > 1 {
		return "", errors.New("multi pir server service found with same service id")
	}
	if resp.Data[0] == nil {
		return "", errors.New("internal error,response Data not exist")
	}
	return resp.Data[0].Message, nil
}

// IsActive 检查服务是否就绪。
func (s *Server) IsActive(ctx context.Context) (bool, error) {
	status, err := s.getStatus(ctx)
	if err != nil {
		return false, err
	}
	if status == enums.PirService_ACTIVE || status == enums.PirService_USING {
		return true, nil
	}
	return false, nil
}

func (s *Server) waitStatus(
	ctx context.Context,
	targetStatus []enums.PirService_Status,
	interval time.Duration,
) (enums.PirService_Status, error) {
	status, err := s.getStatus(ctx)
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
			status, err := s.getStatus(ctx)
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

// WaitForReady 阻塞等待，直到服务就绪。如果检查到部署失败会直接返回错误。
func (s *Server) WaitForReady(ctx context.Context, interval time.Duration) error {
	status, err := s.waitStatus(ctx, serverStableStatus, waitInterval)
	if err != nil {
		return err
	}
	if status != enums.PirService_ACTIVE && status != enums.PirService_USING {
		message, err := s.getMessage(ctx)
		if err != nil {
			return err
		}
		errMsg := fmt.Sprintf("unexpected pir server service status: %s,Message:%s", status.String(), message)
		return errors.New(errMsg)
	}
	return nil
}

// New ClientToken 为客户端参与方创建token。
// token用于部署client端服务。
// 允许为同一party创建多个token，不同token的client端service的pir查询计量会区别统计。
// token创建后默认处于启用状态。
func (s *Server) NewClientToken(ctx context.Context, partyID string) (string, error) {
	createTokenResp, err := s.CreatePirServerToken(ctx, &token.GenerateTokenForPartyRequest{
		PartyId:   partyID,
		ServiceId: s.serviceID,
	})
	if err != nil {
		return "", errors.Wrap(err,
			fmt.Sprintf("pir server:%d service create token for party:%s failed", s.serviceID, partyID))
	}
	return createTokenResp.Token, nil
}

// DisableClientToken 停用token。
func (s *Server) DisableClientToken(ctx context.Context, tokenStr string) error {
	_, err := s.EnablePirServerToken(ctx, &token.ActionOnPirServerToken{
		Token:  tokenStr,
		Action: enums.Token_DISABLE,
	})
	if err != nil {
		return errors.Wrap(err, "disable client token failed")
	}
	return nil
}

// EnableClientToken 启用token。
func (s *Server) EnableClientToken(ctx context.Context, tokenStr string) error {
	_, err := s.EnablePirServerToken(ctx, &token.ActionOnPirServerToken{
		Token:  tokenStr,
		Action: enums.Token_ENABLE,
	})
	if err != nil {
		return errors.Wrap(err, "enable client token failed")
	}
	return nil
}

// 获取serviceID，serviceID是pir server service的属性。
func (s *Server) GetServiceID() uint64 {
	return s.serviceID
}

// Delete 终结并删除service。
func (s *Server) Delete(ctx context.Context) error {
	currentStatus, err := s.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currentStatus != enums.PirService_TERMINATED && currentStatus != enums.PirService_FAILED {
		terminateErr := s.TerminateService(ctx)
		if terminateErr != nil {
			return terminateErr
		}
	}
	err = s.deleteService(ctx)
	if err != nil {
		return err
	}
	return nil
}

// TerminateService 仅终结service。
func (s *Server) TerminateService(ctx context.Context) error {
	_, err := s.PirServerServiceTakeAction(ctx, &pir.ServerServiceTakeActionRequest{
		ServiceIdStr: fmt.Sprintf("%d", s.serviceID),
		Action:       enums.PirService_TERMINATE.String(),
	})
	if err != nil {
		return errors.Wrap(err, "terminate pir server service failed")
	}
	return nil
}

func (s *Server) deleteService(ctx context.Context) error {
	currentStatus, err := s.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currentStatus != enums.PirService_TERMINATED && currentStatus != enums.PirService_FAILED {
		return errors.New("can't delete pir server service,service not terminated or failed")
	}
	_, err = s.DeletePirServerService(ctx, &pir.DeleteServerServiceRequest{
		ServiceIdStr: fmt.Sprintf("%d", s.serviceID),
	})
	if err != nil {
		return errors.Wrap(err, "delete pir server service failed")
	}
	return nil
}

// UpdateServerData 更新server service数据,grpc 接口的简单封装,更新过程中原来的数据正常提供服务。
// 更新状态可通过 [Server.Status] 查询。
func (s *Server) UpdateServerData(ctx context.Context, dataParam *pir.DataModeParams) error {
	resp, err := s.SudoClient.UpdateServerData(ctx, &pir.UpdateServerDataRequest{
		Id:             s.id,
		DataModeParams: dataParam,
	})
	if err != nil {
		return err
	}
	if resp.Status != pir.UpdateServerDataResponse_SUCCESS {
		return status.Error(codes.Internal, resp.Msg)
	}
	return nil
}

// Status 查询并返回状态。
func (s *Server) Status(ctx context.Context) (enums.PirService_Status, error) {
	return s.getStatus(ctx)
}

// ID 返回id属性。
func (s *Server) ID() uint64 {
	return s.id
}

// Name 返回name属性
func (s *Server) Name() string {
	return s.identityName
}

// ServiceID 返回serviceID属性
func (s *Server) ServiceID() uint64 {
	return s.serviceID
}
