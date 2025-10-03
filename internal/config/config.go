package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "db"),
		DBPassword: getEnv("DB_PASSWORD", "db"),
		DBName:     getEnv("DB_NAME", "db"),
	}

	return cfg

}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
