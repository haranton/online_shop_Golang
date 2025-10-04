package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"onlineShop/internal/config"
	"onlineShop/internal/db"
	"onlineShop/internal/handlers"
	"onlineShop/internal/repo"
	"onlineShop/internal/service"
	"testing"
)

func TestCreateUserIntegration(t *testing.T) {
	cfg := config.LoadConfig()
	cfg.DBName = "db_test"

	db.RunMigrations(cfg)

	database := db.GetDB(cfg)
	if database == nil {
		t.Fatal("failed to connect to database")
	}

	if err := database.Exec("DELETE FROM users;").Error; err != nil {
		t.Fatalf("failed to clear users table: %v", err)
	}

	repository := repo.NewReposytory(database)
	services := service.NewService(repository)
	h := handlers.NewHandler(services)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	server := httptest.NewServer(mux)
	defer server.Close()

	user := map[string]string{"login": "Alice", "password": "123"}
	body, _ := json.Marshal(user)

	resp, err := http.Post(server.URL+"/api/users", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	var count int
	err = database.Raw("SELECT COUNT(*) FROM users WHERE login = 'Alice' and password = '123'").Scan(&count).Error
	if err != nil {
		t.Fatalf("failed to check DB: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 record, got %d", count)
	}

}
