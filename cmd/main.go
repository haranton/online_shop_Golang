package main

import (
	"net/http"
	"onlineShop/internal/config"
	"onlineShop/internal/db"
	"onlineShop/internal/handlers"
	"onlineShop/internal/logger"
	"onlineShop/internal/repo"
	"onlineShop/internal/service"
)

func main() {
	// Инициализация логгера (первым делом, так как он нужен везде)
	log := logger.GetLogger("DEBUG")

	// Загрузка конфигурации
	config := config.LoadConfig(log)

	// Обновляем логгер с правильным уровнем из конфига
	log = logger.GetLogger(config.ENV)

	connectDb := db.GetDB(config, log)
	db.RunMigrations(config, log)

	//todo repository, service, handler
	repo := repo.NewReposytory(connectDb, log)
	service := service.NewService(repo)
	handlers := handlers.NewHandler(service, log)

	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux)

	log.Info("Server is running on port", config.AppPort)
	if err := http.ListenAndServe(":"+config.AppPort, mux); err != nil {
		log.Error("Failed to start server:", err)
	}
}
