package db

import (
	"database/sql"
	"fmt"
	"log"
	"onlineShop/internal/config"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("failed to create migration driver:", err)
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+basepath,
		"postgres", driver,
	)
	if err != nil {
		log.Fatal("failed to create migrate instance:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrations:", err)
	}

	log.Println("Database migrations ran successfully")
}
