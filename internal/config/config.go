package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBTimezone string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string

	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	return &Config{
		AppEnv:     Getenv("APP_ENV", "development"),
		AppPort:    Getenv("APP_PORT", "8080"),
		DBHost:     Getenv("DB_HOST", "localhost"),
		DBPort:     Getenv("DB_PORT", "5432"),
		DBUser:     Getenv("DB_USER", "routex"),
		DBPassword: Getenv("DB_PASSWORD", "routex@2026"),
		DBName:     Getenv("DB_NAME", "merchant"),
		DBSSLMode:  Getenv("DB_SSLMODE", "disable"),
		DBTimezone: Getenv("DB_TIMEZONE", "Asia/Ho_Chi_Minh"),

		RedisHost:     Getenv("REDIS_HOST", "localhost"),
		RedisPort:     Getenv("REDIS_PORT", "6379"),
		RedisPassword: Getenv("REDIS_PASSWORD", ""),
		RedisDB:       Getenv("REDIS_DB", "0"),

		JWTSecret: Getenv("JWT_SECRET", "QAAAlVg30jBTW8j2EtOtAEuREhA8nts1i0l4DX4ghKp"),
	}
}

func Getenv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}
