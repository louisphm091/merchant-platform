package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantStatus string

const (
	MerchantStatusPending   MerchantStatus = "PENDING"
	MerchantStatusApproved  MerchantStatus = "APPROVED"
	MerchantStatusRejected  MerchantStatus = "REJECTED"
	MerchantStatusSuspended MerchantStatus = "SUSPENDED"
)

type Merchant struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	MerchantCode string         `gorm:"size:50;uniqueIndex;not null" json:"merchant_code"`
	BusinessName string         `gorm:"size:255;not null" json:"business_name"`
	Email        string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Phone        string         `gorm:"size:20" json:"phone"`
	Status       MerchantStatus `gorm:"size:20;not null;default:PENDING" json:"status"`
	WebhookURL   string         `gorm:"size:500" json:"webhook_url"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (m *Merchant) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
