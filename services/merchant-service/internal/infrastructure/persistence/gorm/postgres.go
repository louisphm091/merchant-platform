package gorm

import (
	"fmt"
	"log"
	"merchant-platform/merchant-service/internal/infrastructure/config"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
		cfg.DBTimezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&model.AdminModel{},
		&model.MerchantModel{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return db
}
