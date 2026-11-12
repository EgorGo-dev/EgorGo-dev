package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort string
	HTTPPort string
}

func Load() *Config {
	_ = godotenv.Load("5_1/.env")
	_ = godotenv.Load(".env")

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Printf("Загружены настройки: gRPC=%s, HTTP=%s", grpcPort, httpPort)
	return &Config{
		GRPCPort: grpcPort,
		HTTPPort: httpPort,
	}
}
