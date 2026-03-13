package router

import (
	"merchant-platform/merchant-service/internal/controller/http/handler"
	"merchant-platform/merchant-service/internal/controller/http/middleware"
	"merchant-platform/merchant-service/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	cfg *config.Config,
	merchantHandler *handler.MerchantHandler,
	adminHandler *handler.AdminHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", merchantHandler.Health)

	api := r.Group("/api")
	{
		merchants := api.Group("/merchants")
		{
			merchants.POST("/register", merchantHandler.Register)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/login", adminHandler.Login)

			adminProtected := admin.Group("")
			adminProtected.Use(middleware.AdminAuthMiddleware(cfg))
			{
				adminProtected.GET("/merchants", adminHandler.ListMerchants)
				adminProtected.POST("/merchants/:id/approve", adminHandler.ApproveMerchant)
			}
		}
	}

	return r
}
