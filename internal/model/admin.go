package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Entity Class

type Admin struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string    `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	FullName  string    `gorm:"size:255;not null" json:"full_name"`
	IsActive  bool      `gorm:"not null;default:false" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
