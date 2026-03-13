package model

import (
	"merchant-platform/merchant-service/internal/domain/auditing"

	"github.com/google/uuid"
)

type MerchantModel struct {
	ID           uuid.UUID         `gorm:"type:uuid;primaryKey"`
	MerchantCode string            `gorm:"size:50;uniqueIndex;not null"`
	BusinessName string            `gorm:"size:255;not null"`
	Email        string            `gorm:"size:255;uniqueIndex;not null"`
	Phone        string            `gorm:"size:20"`
	WebhookURL   string            `gorm:"size:500"`
	Status       string            `gorm:"size:20;not null"`
	Auditing     auditing.Auditing `gorm:"embedded"`
}

func (MerchantModel) TableName() string {
	return "merchants"
}
