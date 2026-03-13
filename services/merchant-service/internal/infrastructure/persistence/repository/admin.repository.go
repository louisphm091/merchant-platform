package repository

import (
	"context"
	"merchant-platform/merchant-service/internal/domain/admin/entity"
	"merchant-platform/merchant-service/internal/domain/admin/repository"
	persistenceModel "merchant-platform/merchant-service/internal/infrastructure/persistence/model"
	"strings"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func (a AdminRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.Admin, error) {

	var model persistenceModel.AdminModel
	err := a.db.WithContext(ctx).Where("email = ?", strings.TrimSpace(strings.ToLower(email))).First(&model).Error

	if err != nil {
		return nil, err
	}

	return entity.Rehydrate(
		model.ID,
		model.Email,
		model.PasswordHash,
		model.FullName,
		model.IsActive,
		model.Auditing,
	), nil
}

func (a AdminRepositoryImpl) Save(ctx context.Context, admin *entity.Admin) error {

	model := persistenceModel.AdminModel{
		ID:           admin.ID(),
		Email:        admin.Email(),
		PasswordHash: admin.PasswordHash(),
		FullName:     admin.FullName(),
		IsActive:     admin.IsActive(),
		Auditing:     admin.Auditing(),
	}

	return a.db.WithContext(ctx).Create(&model).Error
}

func NewAdminRepositoryImpl(db *gorm.DB) repository.AdminRepository {
	return &AdminRepositoryImpl{
		db: db,
	}
}
