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

	paths := []string{".env", "../.env"}
	var err error
	for _, path := range paths {
		err = godotenv.Load(path)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	cfg := &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "db"),
		DBPassword: getEnv("DB_PASSWORD", "db"),
		DBName:     getEnv("DB_NAME", "db"),
	}

	log.Println("Configuration loaded successfully")
	return cfg

}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
