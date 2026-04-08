package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	PORT        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default environments")
	}

	return &Config{
		DB_HOST:     getEnv("DB_HOST", "127.0.0.1"),
		DB_PORT:     getEnv("DB_PORT", "3306"),
		DB_USER:     getEnv("DB_USER", "root"),
		DB_PASSWORD: getEnv("DB_PASSWORD", ""),
		DB_NAME:     getEnv("DB_NAME", "go_db"),
		PORT:        getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
