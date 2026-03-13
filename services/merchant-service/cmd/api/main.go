package main

import (
	"log"
	"merchant-platform/merchant-service/internal/application/command"
	"merchant-platform/merchant-service/internal/application/query"
	"merchant-platform/merchant-service/internal/controller/http/handler"
	"merchant-platform/merchant-service/internal/controller/http/router"
	"merchant-platform/merchant-service/internal/infrastructure/config"
	"merchant-platform/merchant-service/internal/infrastructure/messaging"
	"merchant-platform/merchant-service/internal/infrastructure/persistence"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/gorm"
	"merchant-platform/merchant-service/internal/infrastructure/persistence/repository"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	cfg := config.LoadConfig()

	db := gorm.NewPostgresConnection(cfg)
	persistence.SeedDefaultAdmin(db, cfg)
	kafkaProducer := messaging.NewKafkaProducer(cfg.KafkaBroker)

	defer func() {
		if err := kafkaProducer.Close(); err != nil {
			log.Printf("Failed to close Kafka producer: %v", err)
		}
	}()

	merchantRepo := repository.NewMerchantRepository(db)
	adminRepo := repository.NewAdminRepositoryImpl(db)
	registerMerchantHandler := command.NewRegisterMerchantHandler(merchantRepo, kafkaProducer)
	adminLoginHandler := command.NewAdminLoginHandler(adminRepo, cfg)
	listMerchantsHandler := query.NewListMerchantsHandler(merchantRepo)
	approveMerchantHandler := command.NewApproveMerchantHandler(merchantRepo, kafkaProducer)
	merchantHTTPHandler := handler.NewMerchantHandler(registerMerchantHandler)
	adminHTTPHandler := handler.NewAdminHandler(
		adminLoginHandler,
		listMerchantsHandler,
		approveMerchantHandler)

	r := router.SetupRouter(cfg, merchantHTTPHandler, adminHTTPHandler)

	log.Printf("merchant-service running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start merchant-service: %v", err)
	}
}
