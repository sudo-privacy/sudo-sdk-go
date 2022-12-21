package pir

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sudosdk/pkg/sudoclient"
	"sudosdk/protobuf/basic/protobuf/enums"
	"sudosdk/protobuf/virtualservice/platformpb/pir"
	"sudosdk/protobuf/virtualservice/platformpb/token"
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

type Server struct {
	*sudoclient.SudoClient
	identityName string
	id           uint64
	serviceID    uint64
}

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

func (s *Server) WaitForReady(ctx context.Context, interval time.Duration) error {
	status, err := s.waitStatus(ctx, serverStableStatus, waitInterval)
	if err != nil {
		return err
	}
	if status != enums.PirService_ACTIVE && status != enums.PirService_USING {
		errMsg := fmt.Sprintf("unexpected pir server service status: %s", status.String())
		return errors.New(errMsg)
	}
	return nil
}

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

func (s *Server) GetServiceID() uint64 {
	return s.serviceID
}

func (s *Server) Delete(ctx context.Context) error {
	currStatus, err := s.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currStatus != enums.PirService_TERMINATED && currStatus != enums.PirService_FAILED {
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

// TerminateService 返回后状态立刻变成Terminated,无需等待
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
	currStatus, err := s.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currStatus != enums.PirService_TERMINATED && currStatus != enums.PirService_FAILED {
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
