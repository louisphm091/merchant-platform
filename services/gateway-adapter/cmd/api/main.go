package main

import (
	"log"
	"merchant-platform/gateway-adapter/internal/config"
	"merchant-platform/gateway-adapter/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()
	proxyHandler := handler.NewProxyHandler(cfg)

	r := gin.Default()
	r.GET("/health", proxyHandler.Health)

	api := r.Group("/api")
	{
		merchants := api.Group("/merchants")
		{
			merchants.POST("/register", proxyHandler.ProxyRegisterMerchant)
		}
	}

	log.Printf("api-gateway running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start api-gateway: %v", err)
	}
}
