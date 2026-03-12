package repository

import (
	"github.com/louisphm091/merchant-platform/internal/model"
	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (r *MerchantRepository) Create(merchant *model.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *MerchantRepository) Update(merchant *model.Merchant) error {
	return r.db.Save(merchant).Error
}

func (r *MerchantRepository) FindByEmail(email string) (*model.Merchant, error) {

	var merchant model.Merchant

	err := r.db.Where("email = ?", email).First(&merchant).Error

	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) FindAll() ([]model.Merchant, error) {
	var merchants []model.Merchant
	err := r.db.Find(&merchants).Error
	return merchants, err
}

func (r *MerchantRepository) FindById(id string) (*model.Merchant, error) {

	var merchant model.Merchant

	err := r.db.Where("id = ?", id).First(&merchant).Error

	if err != nil {
		return nil, err
	}

	return &merchant, nil
}
