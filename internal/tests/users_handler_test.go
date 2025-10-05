package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"testing"
)

func TestCreateUserIntegration(t *testing.T) {

	cfg, err := SetupTestEnv()

	if err != nil {
		t.Fatalf("failed to setup env: %v", err)
	}
	defer cfg.TeardownTestEnv()

	user := map[string]string{"login": "Alice", "password": "123"}
	body, _ := json.Marshal(user)

	resp, err := http.Post(cfg.Server.URL+"/api/users", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	var count int
	err = cfg.DB.Raw("SELECT COUNT(*) FROM users WHERE login = 'Alice' and password = '123'").Scan(&count).Error
	if err != nil {
		t.Fatalf("failed to check DB: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 record, got %d", count)
	}

	/// GET /users
	resp, err = http.Get(cfg.Server.URL + "/api/users")
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var users []models.User
	if err := getBodyFromResponse(resp, &users); err != nil {
		t.Fatalf("failed to parse user response: %v", err)
	}

	if len(users) != 1 {
		t.Fatalf("failed to check count users in GET users: %v", err)
	}

	/// GET /users/{id}
	resp, err = http.Get(cfg.Server.URL + "/api/users/1")

	var userResp models.User
	getBodyFromResponse(resp, &userResp)

	if userResp.Login != "Alice" {
		t.Fatalf("expected other login ,got %v", userResp.Login)
	}

}
