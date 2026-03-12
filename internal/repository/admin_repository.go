package repository

import (
	"github.com/louisphm091/merchant-platform/internal/model"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}

func (r *AdminRepository) FindByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *AdminRepository) FindById(id string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
