package repository

import (
	"context"
	"merchant-platform/merchant-service/internal/domain/merchant/entity"
)

type MerchantRepository interface {
	Save(ctx context.Context, merchant *entity.Merchant) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindById(ctx context.Context, id string) (*entity.Merchant, error)
	FindAll(ctx context.Context) ([]*entity.Merchant, error)
	Update(ctx context.Context, merchant *entity.Merchant) error
}
