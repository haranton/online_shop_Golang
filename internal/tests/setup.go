package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"onlineShop/internal/config"
	"onlineShop/internal/db"
	"onlineShop/internal/handlers"
	"onlineShop/internal/logger"
	"onlineShop/internal/repo"
	"onlineShop/internal/service"

	"gorm.io/gorm"
)

type Setup struct {
	DB     *gorm.DB
	Server *httptest.Server
}

func SetupTestEnv() (*Setup, error) {

	logger := logger.GetLogger("DEBUG")
	cfg := config.LoadConfig(logger)
	cfg.DBName = "db_test"
	cfg.DBHost = "db"

	db.RunMigrations(cfg, logger)

	database := db.GetDB(cfg, logger)
	if database == nil {
		return nil, fmt.Errorf("failed to connect to database")
	}

	if err := database.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error; err != nil {
		return nil, fmt.Errorf("failed to clear users table: %v", err)
	}
	if err := database.Exec("TRUNCATE TABLE products RESTART IDENTITY CASCADE").Error; err != nil {
		return nil, fmt.Errorf("failed to clear products table: %v", err)
	}
	if err := database.Exec("TRUNCATE TABLE categories RESTART IDENTITY CASCADE").Error; err != nil {
		return nil, fmt.Errorf("failed to clear categories table: %v", err)
	}
	if err := database.Exec("TRUNCATE TABLE orders RESTART IDENTITY CASCADE").Error; err != nil {
		return nil, fmt.Errorf("failed to clear orders table: %v", err)
	}

	if err := database.Exec("TRUNCATE TABLE order_products RESTART IDENTITY CASCADE").Error; err != nil {
		return nil, fmt.Errorf("failed to clear order_products table: %v", err)
	}

	repository := repo.NewReposytory(database, logger)
	services := service.NewService(repository)
	h := handlers.NewHandler(services, logger)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	server := httptest.NewServer(mux)

	return &Setup{
		DB:     database,
		Server: server,
	}, nil

}

func (s *Setup) TeardownTestEnv() {

	s.Server.Close()
}
