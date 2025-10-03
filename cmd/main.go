package main

import (
	"onlineShop/internal/db"
)

func main() {
	// database and migration
	connectDb := db.GetDB()
	db.RunMigrations()

	_ = connectDb

	//todo repository, service, handler
}
