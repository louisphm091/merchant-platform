package repository

import (
	"context"
	"merchant-platform/merchant-service/internal/domain/admin/entity"
)

type AdminRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.Admin, error)
	Save(ctx context.Context, admin *entity.Admin) error
}
