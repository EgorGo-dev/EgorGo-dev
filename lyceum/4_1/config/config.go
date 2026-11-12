package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort string
}

func Load() *Config {
	err := godotenv.Load("4_1/.env")
	if err != nil {
		_ = godotenv.Load(".env")
	}

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		log.Println("GRPC_PORT не задан, используется порт по умолчанию 50051")
		port = "50051"
	}

	return &Config{
		GRPCPort: port,
	}
}
