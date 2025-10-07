package main

import (
	"net/http"
	"onlineShop/internal/config"
	"onlineShop/internal/db"
	"onlineShop/internal/handlers"
	"onlineShop/internal/logger"
	"onlineShop/internal/repo"
	"onlineShop/internal/service"

	"github.com/rs/cors"
)

func main() {
	// Инициализация логгера (первым делом, так как он нужен везде)
	log := logger.GetLogger("DEBUG")

	// Загрузка конфигурации
	cfg := config.LoadConfig(log)

	// Обновляем логгер с правильным уровнем из конфига
	log = logger.GetLogger(cfg.ENV)

	connectDb := db.GetDB(cfg, log)
	db.RunMigrations(cfg, log)

	// Инициализация слоев приложения
	repository := repo.NewReposytory(connectDb, log)
	services := service.NewService(repository)
	handler := handlers.NewHandler(services, log)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	// Настройка CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // React dev server
			"http://localhost:5173", // Vite dev server
			"http://localhost:8080", // Ваш backend (если нужно)
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH",
		},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "Content-Length",
			"Authorization", "X-CSRF-Token", "X-Requested-With",
		},
		AllowCredentials: true,
		Debug:            cfg.ENV == "DEBUG", // Логи только в режиме отладки
	})

	// Обертываем mux в CORS handler
	handlerWithCORS := corsHandler.Handler(mux)

	log.Info("Server is running on port", "port", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, handlerWithCORS); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
