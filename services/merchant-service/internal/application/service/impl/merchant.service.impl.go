package impl

import (
	"context"
	"merchant-platform/merchant-service/internal/application/dto"
	"merchant-platform/merchant-service/internal/application/service"
	"merchant-platform/merchant-service/internal/domain/merchant/entity"
)

type merchantService struct {
}

func (m merchantService) Register(ctx context.Context, request *dto.RegisterMerchantRequest) (*entity.Merchant, error) {
	//TODO implement me
	panic("implement me")
}

func (m merchantService) List(ctx context.Context) ([]entity.Merchant, error) {
	//TODO implement me
	panic("implement me")
}

func (m merchantService) ApproveMerchant(ctx context.Context, id string) (*entity.Merchant, error) {
	//TODO implement me
	panic("implement me")
}

func (m merchantService) RejectMerchant(ctx context.Context, id string) (*entity.Merchant, error) {
	//TODO implement me
	panic("implement me")
}

func NewMerchantService() service.MerchantService {
	return &merchantService{}
}
