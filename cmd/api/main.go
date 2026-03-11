package main

import (
	"log"

	"github.com/louisphm091/merchant-platform/internal/cache"
	"github.com/louisphm091/merchant-platform/internal/config"
	"github.com/louisphm091/merchant-platform/internal/database"
	"github.com/louisphm091/merchant-platform/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	db := database.NewPostgresConnection(cfg)
	database.AutoMigrate(db)

	rdb := cache.NewRedisConnection(cfg)

	r := router.SetupRouter(cfg, db, rdb)

	log.Printf("Server is running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start server : %v", err)
	}
}
