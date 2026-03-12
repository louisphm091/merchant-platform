package database

import (
	"log"
	"strings"

	"github.com/louisphm091/merchant-platform/internal/model"
	"github.com/louisphm091/merchant-platform/internal/utils"
	"gorm.io/gorm"
)

func SeedAdmins(db *gorm.DB) {
	const defaultAdminEmail = "admin@merchant-platform.local"
	const defaultAdminPassword = "admin123"
	const defaultAdminName = "System Admin"

	var existing model.Admin
	err := db.Where("email = ?", defaultAdminEmail).First(&existing).Error
	if err == nil {
		log.Println("Default admin already exists")
		return
	}

	if err != gorm.ErrRecordNotFound && err != nil {
		log.Printf("failed to check default admin: %v", err)
		return
	}

	hashedPassword, err := utils.HashPassword(defaultAdminPassword)
	if err != nil {
		log.Printf("failed to hash default admin password: %v", err)
		return
	}

	admin := model.Admin{
		Email:    strings.ToLower(defaultAdminEmail),
		Password: hashedPassword,
		FullName: defaultAdminName,
		IsActive: true,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("failed to seed default admin: %v", err)
		return
	}

	log.Println("Default admin seeded successfully")
	log.Println("email: admin@merchant-platform.local")
	log.Println("password: admin123")
}
