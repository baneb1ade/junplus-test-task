package config

import (
	"log"
	"os"
)

type Config struct {
	Server  ServerConfig
	Storage StorageConfig
}

type StorageConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

type ServerConfig struct {
	Address string
	Port    string
}

func MustLoad() *Config {

	requiredEnvVars := []string{
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_DB",
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Missing required environment variable: %s", envVar)
		}
	}

	config := Config{
		Server: ServerConfig{Address: os.Getenv("SERVER_ADDRESS"), Port: os.Getenv("SERVER_PORT")},
		Storage: StorageConfig{
			DBUser:     os.Getenv("POSTGRES_USER"),
			DBPassword: os.Getenv("POSTGRES_PASSWORD"),
			DBHost:     os.Getenv("POSTGRES_HOST"),
			DBPort:     os.Getenv("POSTGRES_PORT"),
			DBName:     os.Getenv("POSTGRES_DB"),
		},
	}

	return &config
}
