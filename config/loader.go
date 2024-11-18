package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return &Config{
		Server{
			Host: os.Getenv("HOST"),
			Port: os.Getenv("PORT"),
		},
		Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT{
			AccessSecretKey:  os.Getenv("ACCESS_TOKEN_KEY"),
			RefreshSecretKey: os.Getenv("REFRESH_TOKEN_KEY"),
		},
	}
}
