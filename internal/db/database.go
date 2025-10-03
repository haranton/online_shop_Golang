package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"

	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("Database connection is successful")
	return db
}

func RunMigrations() {
	dsn := "postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	driver, err := migratePostgres.WithInstance(db, &migratePostgres.Config{})
	if err != nil {
		log.Fatal("failed to create migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
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
