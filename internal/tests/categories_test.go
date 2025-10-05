package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"testing"
)

func TestCategoriesIntegration(t *testing.T) {

	cfg, err := SetupTestEnv()

	if err != nil {
		t.Fatalf("failed to setup env: %v", err)
	}
	defer cfg.TeardownTestEnv()

	resp, err := http.Get(cfg.Server.URL + "/api/categories")
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var categories []map[string]any
	if err := getBodyFromResponse(resp, &categories); err != nil {
		t.Fatalf("failed to parse categories response: %v", err)
	}

	if len(categories) != 0 {
		t.Fatalf("expected 0 categories, got %d", len(categories))
	}

	//post category
	categoriesPost := []map[string]string{
		{"name": "Electronics"},
		{"name": "Books"},
		{"name": "Books1"},
		{"name": ""},
		{"name": "test"},
	}

	for _, cat := range categoriesPost {

		body, _ := json.Marshal(cat)

		resp, err := http.Post(cfg.Server.URL+"/api/categories", "application/json", bytes.NewReader(body))
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", resp.StatusCode)
		}

	}

	var categoriesDB []models.Category
	err = cfg.DB.Raw("SELECT * FROM categories").Scan(&categoriesDB).Error
	if err != nil {
		t.Fatalf("failed to query db")
	}

	if len(categoriesDB) != 5 {
		t.Fatalf("expected 5 categories, got %d", len(categoriesDB))
	}
}
