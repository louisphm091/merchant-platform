package cache

import (
	"context"
	"fmt"
	"log"
	"merchant-platform/merchant-service/internal/infrastructure/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(cfg *config.Config) *redis.Client {

	dbIndex, err := strconv.Atoi(cfg.RedisDB)

	if err != nil {
		dbIndex = 0
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       dbIndex,
	})

	_, err = client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis Successfully")
	return client
}
