package onlinesvc

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/apl"
	"github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/online_service"
)

var waitInterval = time.Second * 5

var serviceStableStatus = []enums.Service_Status{
	enums.Service_UNKNOWN_STATUS,
	enums.Service_OPEN,
	enums.Service_CANCELED,
	enums.Service_DENIED,
	enums.Service_IDLE,
	enums.Service_ACTIVE,
	enums.Service_FAILED,
	enums.Service_TERMINATED,
}

type Service struct {
	*Factory
	*online_service.Service
}

// Predict 数据使用方使用该方法预测请求。
func (service *Service) Predict(
	ctx context.Context,
	req *online_service.PredictRequest,
) (*online_service.PredictResponse, error) {
	if service.Base == nil {
		return nil, errors.New("service.Base is nil")
	}
	if service.Base.OwnerPartyId != service.Factory.PartyID {
		return nil, errors.New("you are not owner of this Service")
	}
	splitPath := strings.Split(service.Base.Path, "/")
	req.SubPath = splitPath[len(splitPath)-1]
	return service.Factory.SudoClient.OnlinePredict(ctx, req)
}

// Approve 数据提供方许可服务申请。
func (service *Service) Approve(ctx context.Context) error {
	if service.Base == nil {
		return errors.New("service.Base is nil")
	}
	if service.Base.OwnerPartyId == service.Factory.PartyID {
		return errors.New("can not approve self service")
	}
	_, err := service.Factory.SudoClient.ApplicationTakeAction(ctx, &apl.AplTakeActionRequest{
		Param: &apl.AplTakeActionParam{
			Action: enums.Application_APPROVE,
			RefId:  service.Base.ServiceId,
			Type:   "predictService",
		},
		ToTusita: service.Factory.TusitaID,
	})
	if err != nil {
		return errors.Wrapf(err, "falied to approve service %s", service.Base.Name)
	}
	return nil
}

// Reject 数据提供方拒绝服务申请。
func (service *Service) Reject(ctx context.Context) error {
	if service.Base == nil {
		return errors.New("service.Base is nil")
	}
	if service.Base.OwnerPartyId != service.Factory.PartyID {
		return errors.New("can not reject self service")
	}
	_, err := service.Factory.SudoClient.ApplicationTakeAction(ctx, &apl.AplTakeActionRequest{
		Param: &apl.AplTakeActionParam{
			Action: enums.Application_REJECT,
			RefId:  service.Base.ServiceId,
			Type:   "predictService",
		},
		ToTusita: service.Factory.TusitaID,
	})
	if err != nil {
		return errors.Wrapf(err, "falied to reject service %s", service.Base.Name)
	}
	return nil
}

// CancelService 服务发起方 撤销申请。
func (service *Service) CancelService(ctx context.Context) error {
	_, err := service.Factory.SudoClient.ServiceTakeAction(ctx, &online_service.ServiceTakeActionRequest{
		Param: &online_service.ServiceTakeActionParam{
			ServiceId: service.Base.ServiceId,
			Action:    enums.Service_CANCEL,
		},
		ToTusita: service.Factory.TusitaID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to terminate Service")
	}
	return nil
}

// ReapplyService 服务发起方在数据提供方拒绝申请后重新申请。
func (service *Service) ReapplyService(ctx context.Context) error {
	_, err := service.Factory.SudoClient.ServiceTakeAction(ctx, &online_service.ServiceTakeActionRequest{
		Param: &online_service.ServiceTakeActionParam{
			ServiceId: service.Base.ServiceId,
			Action:    enums.Service_REAPPLY,
		},
		ToTusita: service.Factory.TusitaID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to reapply Service")
	}
	return nil
}

// TerminateService 停止服务，服务发起方和数据提供方均可操作。
func (service *Service) TerminateService(ctx context.Context) error {
	_, err := service.Factory.SudoClient.ServiceTakeAction(ctx, &online_service.ServiceTakeActionRequest{
		Param: &online_service.ServiceTakeActionParam{
			ServiceId: service.Base.ServiceId,
			Action:    enums.Service_TERMINATE,
		},
		ToTusita: service.Factory.TusitaID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to terminate Service")
	}
	return nil
}

// RebootService 停止服务后，重启服务。服务发起方和数据提供方均可操作。
func (service *Service) RebootService(ctx context.Context) error {
	_, err := service.Factory.SudoClient.ServiceTakeAction(ctx, &online_service.ServiceTakeActionRequest{
		Param: &online_service.ServiceTakeActionParam{
			ServiceId: service.Base.ServiceId,
			Action:    enums.Service_REBOOT,
		},
		ToTusita: service.Factory.TusitaID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to reboot Service")
	}
	return nil
}

// Delete 删除服务。
func (service *Service) Delete(ctx context.Context) error {
	currentStatus, err := service.getStatus(ctx)
	if err != nil {
		sourceErr := errors.Cause(err)
		if status, ok := status.FromError(sourceErr); ok {
			if status.Code() == codes.NotFound {
				return nil
			}
		}
	}
	if currentStatus == enums.Service_BOOTING ||
		currentStatus == enums.Service_ACTIVE {
		tErr := service.TerminateService(ctx)
		if tErr != nil {
			return errors.Wrap(tErr, "terminate pir client service failed")
		}
		wErr := service.WaitForTerminate(ctx)
		if wErr != nil {
			return wErr
		}
	}

	if currentStatus == enums.Service_OPEN {
		err = service.CancelService(ctx)
		if err != nil {
			return err
		}
	}
	err = service.deleteService(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) deleteService(ctx context.Context) error {
	_, err := service.Factory.SudoClient.DeleteService(ctx, &online_service.DeleteServiceRequest{
		ServiceId: service.Base.ServiceId,
		ToTusita:  service.Factory.TusitaID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to delete service")
	}
	return nil
}

// WaitForReady 等待服务部署完成。
func (service *Service) WaitForReady(ctx context.Context) error {
	status, err := service.waitStatus(ctx, serviceStableStatus, waitInterval)
	if err != nil {
		return err
	}
	if status != enums.Service_ACTIVE {
		message, err := service.getMessage(ctx)
		if err != nil {
			return err
		}
		errMsg := fmt.Sprintf("unexpected pir client service status: %s,Message: %s", status.String(), message)
		return errors.New(errMsg)
	}
	return nil
}

// WaitForTerminate 等待服务停止完成。
func (service *Service) WaitForTerminate(ctx context.Context) error {
	status, err := service.waitStatus(ctx, serviceStableStatus, waitInterval)
	if err != nil {
		return err
	}
	if status != enums.Service_TERMINATED {
		message, err := service.getMessage(ctx)
		if err != nil {
			return err
		}
		errMsg := fmt.Sprintf("unexpected pir client service status: %s,Message: %s", status.String(), message)
		return errors.New(errMsg)
	}
	return nil
}

func (service *Service) waitStatus(
	ctx context.Context,
	targetStatus []enums.Service_Status,
	interval time.Duration,
) (enums.Service_Status, error) {
	status, err := service.getStatus(ctx)
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
			status, err := service.getStatus(ctx)
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

func (service *Service) getStatus(ctx context.Context) (enums.Service_Status, error) {
	tab := enums.Service_SELF
	if service.Base.OwnerPartyId != service.Factory.PartyID {
		tab = enums.Service_OTHERS
	}
	resp, err := service.Factory.SudoClient.GetServices(ctx, &online_service.GetServicesRequest{
		Filter: &online_service.ServiceFilter{
			ServiceIds: []uint64{service.Base.ServiceId},
		},
		Tab:      tab.String(),
		ToTusita: service.Factory.TusitaID,
	})
	if err != nil {
		return 0, errors.Wrap(err, "get service failed")
	}
	if len(resp.Data) == 0 {
		notFoundErr := status.Error(codes.NotFound, "service not found")
		return 0, errors.Wrap(notFoundErr, "no data in get service response")
	}
	if len(resp.Data) > 1 {
		return 0, errors.New("multi service found with same service id")
	}
	if resp.Data[0] == nil {
		return 0, errors.Wrap(err, "get nil service")
	}
	statusNum := enums.Service_Status_value[resp.Data[0].Status]
	if statusNum == 0 {
		return 0, errors.Errorf("unknown service status %s", resp.Data[0].Status)
	}
	return enums.Service_Status(statusNum), nil
}

func (service *Service) getMessage(ctx context.Context) (string, error) {
	resp, err := service.Factory.SudoClient.GetService(ctx, &online_service.GetServiceRequest{
		ServiceId: strconv.FormatUint(service.Base.ServiceId, 10),
		ToTusita:  service.Factory.TusitaID,
	})
	if err != nil {
		return "", errors.Wrap(err, "get service failed")
	}
	if resp.Data == nil || resp.Data.Service == nil || resp.Data.Service.Base == nil {
		return "", errors.Wrap(err, "get nil service")
	}
	return resp.Data.Service.Base.Message, nil
}

// ID 返回 serviceID属性
func (service *Service) ID() uint64 {
	return service.Base.ServiceId
}

// Name 返回 Name属性
func (service *Service) Name() string {
	return service.Base.Name
}

// OwnerPartyID 返回OwnerPartyID属性
func (service *Service) OwnerPartyID() string {
	return service.Base.OwnerPartyId
}
