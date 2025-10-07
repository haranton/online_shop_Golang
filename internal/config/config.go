package config

import (
	"log/slog"
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
	ENV        string
}

// LoadConfig принимает логгер как параметр
func LoadConfig(logger *slog.Logger) *Config {
	paths := []string{".env", "../.env"}
	var err error
	for _, path := range paths {
		err = godotenv.Load(path)
		if err == nil {
			logger.Debug("Loaded environment file", "path", path)
			break
		}
	}
	if err != nil {
		logger.Warn(".env file not found, using system environment variables")
	}

	cfg := &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "db"),
		DBPassword: getEnv("DB_PASSWORD", "db"),
		DBName:     getEnv("DB_NAME", "db"),
		ENV:        getEnv("ENV", "DEBUG"),
	}

	logger.Info("Configuration loaded successfully",
		"app_port", cfg.AppPort,
		"db_host", cfg.DBHost,
		"db_name", cfg.DBName,
		"env", cfg.ENV,
	)

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
