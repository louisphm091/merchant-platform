package service

import (
	"context"
	"merchant-platform/merchant-service/internal/application/dto"
	"merchant-platform/merchant-service/internal/domain/merchant/entity"
)

type MerchantService interface {
	Register(ctx context.Context, request *dto.RegisterMerchantRequest) (*entity.Merchant, error)
	List(ctx context.Context) ([]entity.Merchant, error)
	ApproveMerchant(ctx context.Context, id string) (*entity.Merchant, error)
	RejectMerchant(ctx context.Context, id string) (*entity.Merchant, error)
}
