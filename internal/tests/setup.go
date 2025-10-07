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

	// Сначала подключаемся к основной базе для создания тестовой
	cfg.DBName = "postgres" // Подключаемся к системной базе
	cfg.DBHost = "db"

	// Создаем подключение к системной базе данных
	systemDB := db.GetDB(cfg, logger)
	if systemDB == nil {
		return nil, fmt.Errorf("failed to connect to system database")
	}

	// Создаем тестовую базу если не существует
	if err := createTestDatabase(systemDB, "db_test"); err != nil {
		return nil, fmt.Errorf("failed to create test database: %v", err)
	}

	// Закрываем подключение к системной базе
	sqlDB, _ := systemDB.DB()
	sqlDB.Close()

	// Теперь подключаемся к тестовой базе
	cfg.DBName = "db_test"
	testDB := db.GetDB(cfg, logger)
	if testDB == nil {
		return nil, fmt.Errorf("failed to connect to test database")
	}

	// Запускаем миграции на тестовой базе
	db.RunMigrations(cfg, logger)

	// Очищаем таблицы
	if err := clearAllTables(testDB); err != nil {
		return nil, fmt.Errorf("failed to clear tables: %v", err)
	}

	repository := repo.NewReposytory(testDB, logger)
	services := service.NewService(repository)
	h := handlers.NewHandler(services, logger)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	server := httptest.NewServer(mux)

	return &Setup{
		DB:     testDB,
		Server: server,
	}, nil
}

// createTestDatabase создает тестовую базу данных если она не существует
func createTestDatabase(db *gorm.DB, dbName string) error {
	// Проверяем существует ли база
	var count int64
	result := db.Raw(`
		SELECT COUNT(*) 
		FROM pg_database 
		WHERE datname = ?
	`, dbName).Scan(&count)

	if result.Error != nil {
		return fmt.Errorf("failed to check database existence: %v", result.Error)
	}

	// Если база не существует - создаем
	if count == 0 {
		createSQL := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if err := db.Exec(createSQL).Error; err != nil {
			return fmt.Errorf("failed to create database %s: %v", dbName, err)
		}
		logger.GetLogger("DEBUG").Info("Test database created", "database", dbName)
	} else {
		logger.GetLogger("DEBUG").Info("Test database already exists", "database", dbName)
	}

	return nil
}

// clearAllTables очищает все таблицы в тестовой базе
func clearAllTables(db *gorm.DB) error {
	tables := []string{
		"order_products",
		"orders",
		"products",
		"categories",
		"users",
	}

	for _, table := range tables {
		// Проверяем существует ли таблица перед очисткой
		var tableExists bool
		db.Raw(`
			SELECT EXISTS (
				SELECT FROM information_schema.tables 
				WHERE table_schema = 'public' 
				AND table_name = ?
			)
		`, table).Scan(&tableExists)

		if tableExists {
			if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)).Error; err != nil {
				return fmt.Errorf("failed to clear table %s: %v", table, err)
			}
		}
	}

	return nil
}

func (s *Setup) TeardownTestEnv() {
	// Закрываем подключение к базе
	if s.DB != nil {
		sqlDB, _ := s.DB.DB()
		sqlDB.Close()
	}

	// Закрываем тестовый сервер
	s.Server.Close()
}
