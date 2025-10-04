package main

import (
	"fmt"
	"net/http"
	"onlineShop/internal/config"
	"onlineShop/internal/db"
	"onlineShop/internal/handlers"
	"onlineShop/internal/repo"
	"onlineShop/internal/service"
)

func main() {
	// database and migration
	config := config.LoadConfig()
	connectDb := db.GetDB(config)
	db.RunMigrations(config)

	_ = connectDb

	//todo repository, service, handler
	repo := repo.NewReposytory(connectDb)
	service := service.NewService(repo)
	handlers := handlers.NewHandler(service)

	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux)

	fmt.Println("Server is running on port", config.AppPort)
	if err := http.ListenAndServe(":"+config.AppPort, mux); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
