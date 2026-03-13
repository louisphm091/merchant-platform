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

	KafkaBroker string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string

	AdminEmail    string
	AdminPassword string
	AdminName     string

	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	return &Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "routex"),
		DBPassword: getEnv("DB_PASSWORD", "routex@2026"),
		DBName:     getEnv("DB_NAME", "merchant"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		DBTimezone: getEnv("DB_TIMEZONE", "Asia/Ho_Chi_Minh"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnv("REDIS_DB", "0"),

		AdminEmail:    getEnv("ADMIN_EMAIL", "admin@merchant-routex.local"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "routex@2026"),
		AdminName:     getEnv("ADMIN_NAME", "routex"),
		KafkaBroker:   getEnv("KAFKA_BROKER", "localhost:9092"),

		JWTSecret: getEnv("JWT_SECRET", "QAAAlVg30jBTW8j2EtOtAEuREhA8nts1i0l4DX4ghKp"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
