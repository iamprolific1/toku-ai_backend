package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)


type Config struct {
	Port string
	DBHost string
	DBUser string
	DBPass string
	DBPort string
	DBName string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment")
	}

	cfg := &Config{
		Port:           getEnv("PORT", "8080"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPass:     getEnv("DB_PASSWORD", "password"),
		DBName:         getEnv("DB_NAME", "Tokuai"),
		
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}