package model

import (
	"merchant-platform/merchant-service/internal/domain/auditing"

	"github.com/google/uuid"
)

type AdminModel struct {
	ID           uuid.UUID         `gorm:"type:uuid;primaryKey"`
	Email        string            `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash string            `gorm:"type:varchar(255);not null"`
	FullName     string            `gorm:"type:varchar(255);not null"`
	IsActive     bool              `gorm:"type:boolean;not null;default:true"`
	Auditing     auditing.Auditing `gorm:"embedded"`
}

func (AdminModel) TableName() string {
	return "admin"
}
