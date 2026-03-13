package config

import "os"

type Config struct {
	AppPort            string
	MerchantServiceURL string
}

func Load() *Config {
	return &Config{
		AppPort:            getEnv("APP_PORT", "8080"),
		MerchantServiceURL: getEnv("MERCHANT_SERVICE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}
