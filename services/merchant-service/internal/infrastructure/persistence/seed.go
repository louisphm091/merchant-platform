package persistence

import (
	"context"
	"errors"
	"log"
	"merchant-platform/merchant-service/internal/domain/admin/entity"
	"merchant-platform/merchant-service/internal/infrastructure/config"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/repository"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/utils"

	"gorm.io/gorm"
)

func SeedDefaultAdmin(db *gorm.DB, cfg *config.Config) {
	repo := repository.NewAdminRepositoryImpl(db)

	existing, err := repo.FindByEmail(context.Background(), cfg.AdminEmail)
	if err == nil && existing != nil {
		log.Println("default admin already exists")
		return
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("failed to check default admin: %v", err)
		return
	}

	hashedPassword, err := utils.HashPassword(cfg.AdminPassword)
	if err != nil {
		log.Printf("failed to hash admin password: %v", err)
		return
	}

	admin, err := entity.NewAdmin(cfg.AdminEmail, hashedPassword, cfg.AdminName)
	if err != nil {
		log.Printf("failed to create default admin entity: %v", err)
		return
	}

	if err := repo.Save(context.Background(), admin); err != nil {
		log.Printf("failed to save default admin: %v", err)
		return
	}

	log.Println("default admin seeded successfully")
}
