package main

import (
	"fmt"
	"net/http"
	"onlineShop/internal/config"
	"onlineShop/internal/db"
)

func main() {
	// database and migration
	config := config.LoadConfig()
	connectDb := db.GetDB(config)
	db.RunMigrations(config)

	_ = connectDb

	//todo repository, service, handler

	// server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println("Server is running on port", config.AppPort)
	if err := http.ListenAndServe(":"+config.AppPort, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
