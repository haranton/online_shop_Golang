package config

import (
	"log/slog"
	"os"
	"path/filepath"

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

// LoadConfig загружает конфигурацию из .env или переменных окружения
func LoadConfig(logger *slog.Logger) *Config {
	// Определяем, где находимся
	cwd, _ := os.Getwd()
	logger.Debug("Current working directory", "cwd", cwd)

	// Возможные пути для поиска .env (покрывает тесты и прод)
	paths := []string{
		".env",
		"../.env",
		"../../.env",
		"/app/.env", // на случай, если контейнер в Docker-контексте
	}

	var loadedPath string
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			if err := godotenv.Load(path); err == nil {
				loadedPath, _ = filepath.Abs(path)
				logger.Info("Loaded .env file", "path", loadedPath)
				break
			}
		}
	}

	if loadedPath == "" {
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
