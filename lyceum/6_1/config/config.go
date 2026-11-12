package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort string
	HTTPPort string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

func Load() *Config {
	_ = godotenv.Load("6_1/.env")
	_ = godotenv.Load(".env")

	return &Config{
		GRPCPort: getEnv("GRPC_PORT", "50051"),
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		DBHost:   getEnv("POSTGRES_HOST", "localhost"),
		DBPort:   getEnv("POSTGRES_PORT", "5432"),
		DBUser:   getEnv("POSTGRES_USER", "postgres"),
		DBPass:   getEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:   getEnv("POSTGRES_DB", "orders"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
