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

	if len(categoriesDB) != 4 {
		t.Fatalf("expected 4 categories, got %d", len(categoriesDB))
	}

	//get category by id
	resp, err = http.Get(cfg.Server.URL + "/api/categories/1")
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var category models.Category
	if err := getBodyFromResponse(resp, &category); err != nil {
		t.Fatalf("failed to parse category response: %v", err)
	}

	if len(category.Name) == 0 {
		t.Fatalf("expected category name, got empty")
	}

	//update category
	updateCat := map[string]any{"id": 1, "name": "Updated Electronics"}
	body, _ := json.Marshal(updateCat)

	req, err := http.NewRequest(http.MethodPut, cfg.Server.URL+"/api/categories/1", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var updatedCategory models.Category
	if err := getBodyFromResponse(resp, &updatedCategory); err != nil {
		t.Fatalf("failed to parse updated category response: %v", err)
	}

	if updatedCategory.Name != "Updated Electronics" {
		t.Fatalf("expected updated category name, got %s", updatedCategory.Name)
	}
}
