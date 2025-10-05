package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"testing"
)

func TestProducts(t *testing.T) {

	cfg, err := SetupTestEnv()
	if err != nil {
		t.Fatalf("failed load env %v", err)
	}

	t.Log("products tests")
	// test create product

	var products []map[string]any
	products = append(products, map[string]any{
		"name":       "43242",
		"count":      10,
		"CategoryID": 1,
	})
	products = append(products, map[string]any{
		"name":       "Smartphone",
		"count":      10,
		"CategoryID": 1,
	})
	products = append(products, map[string]any{
		"name":       "Book",
		"count":      10,
		"CategoryID": 1,
	})

	categoriesPost := map[string]string{"name": "Electronics"}

	body, _ := json.Marshal(categoriesPost)

	resp, err := http.Post(cfg.Server.URL+"/api/categories", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	for _, product := range products {
		productJSON, _ := json.Marshal(product)
		resp, err := http.Post(cfg.Server.URL+"/api/products", "application/json", bytes.NewReader(productJSON))
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected status 201, has got : %v", resp.StatusCode)
		}

	}

	var productsDB []models.Product
	if err := cfg.DB.Raw("SELECT * from products").Scan(&productsDB).Error; err != nil {
		t.Fatalf("failed to query to db err: %v", err)
	}

	if len(products) != 3 {
		t.Fatalf("expected 3 products, has got : %v", len(productsDB))
	}

}
