package router

import (
	"github.com/gin-gonic/gin"
	"github.com/louisphm091/merchant-platform/internal/config"
	"github.com/louisphm091/merchant-platform/internal/handler"
	"github.com/louisphm091/merchant-platform/internal/repository"
	"github.com/louisphm091/merchant-platform/internal/service"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB, rdb *redis.Client) *gin.Engine {

	r := gin.Default()

	healthHandler := handler.NewHealthHandler()

	merchantRepo := repository.NewMerchantRepository(db)
	merchantService := service.NewMerchantService(merchantRepo)
	merchantHandler := handler.NewMerchantHandler(merchantService)

	api := r.Group("/api")
	{
		api.GET("/health", healthHandler.HealthCheck)

		merchant := api.Group("/merchants")
		{
			merchant.POST("/register", merchantHandler.Register)
			merchant.GET("", merchantHandler.List)
		}
	}

	return r
}
