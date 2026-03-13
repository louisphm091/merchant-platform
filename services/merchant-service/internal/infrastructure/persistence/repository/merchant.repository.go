package repository

import (
	"context"
	"merchant-platform/merchant-service/internal/domain/merchant/entity"
	"merchant-platform/merchant-service/internal/domain/merchant/repository"
	"merchant-platform/merchant-service/internal/domain/merchant/valueobject"
	persistenceModel "merchant-platform/merchant-service/internal/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func (r *merchantRepository) Update(ctx context.Context, merchant *entity.Merchant) error {

	return r.db.WithContext(ctx).
		Model(&persistenceModel.MerchantModel{}).Where("id = ?", merchant.ID()).Updates(map[string]interface{}{
		"status":     string(merchant.Status()),
		"updated_at": merchant.Auditing().UpdatedAt,
	}).Error
}

func (r *merchantRepository) FindAll(ctx context.Context) ([]*entity.Merchant, error) {
	var models []persistenceModel.MerchantModel

	if err := r.db.WithContext(ctx).Order("created_at desc").Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.Merchant, 0, len(models))

	for _, model := range models {
		merchant, err := toDomain(model)

		if err != nil {
			return nil, err
		}

		result = append(result, merchant)
	}

	return result, nil

}

func (r *merchantRepository) FindById(ctx context.Context, id string) (*entity.Merchant, error) {
	var model persistenceModel.MerchantModel

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}

	return toDomain(model)
}

func (r *merchantRepository) Save(ctx context.Context, merchant *entity.Merchant) error {

	model := persistenceModel.MerchantModel{
		ID:           merchant.ID(),
		MerchantCode: merchant.MerchantCode(),
		BusinessName: merchant.BusinessName(),
		Email:        merchant.Email(),
		Phone:        merchant.Phone(),
		WebhookURL:   merchant.WebhookURL(),
		Status:       string(merchant.Status()),
		Auditing:     merchant.Auditing(),
	}

	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *merchantRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).Model(persistenceModel.MerchantModel{}).Where("email = ?", email).Count(&count).Error

	return count > 0, err
}

func toDomain(model persistenceModel.MerchantModel) (*entity.Merchant, error) {
	email, err := valueobject.NewEmail(model.Email)

	if err != nil {
		return nil, err
	}

	return entity.Rehydrate(
		uuid.MustParse(model.ID.String()),
		model.MerchantCode,
		model.BusinessName,
		email,
		model.Phone,
		model.WebhookURL,
		entity.MerchantStatus(model.Status),
		model.Auditing,
	), nil
}

func NewMerchantRepository(db *gorm.DB) repository.MerchantRepository {
	return &merchantRepository{db: db}
}
