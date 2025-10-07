package db

import (
	"fmt"
	"log/slog"
	"onlineShop/internal/config"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetDB(cfg *config.Config, logger *slog.Logger) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database",
			"error", err,
			"host", cfg.DBHost,
			"port", cfg.DBPort,
			"dbname", cfg.DBName,
		)
		os.Exit(1)
	}

	logger.Info("Database connection is successful")
	return db
}
