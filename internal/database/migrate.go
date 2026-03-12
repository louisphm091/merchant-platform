package database

import (
	"log"

	"github.com/louisphm091/merchant-platform/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Merchant{},
		&model.Admin{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")
}
