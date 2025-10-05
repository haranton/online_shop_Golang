package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"testing"
)

func TestOrders(t *testing.T) {
	cfg, err := SetupTestEnv()
	if err != nil {
		t.Fatalf("failed to load env: %v", err)
	}

	t.Log("orders tests")

	// === 1. Создаём категорию ===
	category := map[string]string{"name": "Electronics"}
	body, _ := json.Marshal(category)
	resp, err := http.Post(cfg.Server.URL+"/api/categories", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	// === 2. Создаём продукты ===
	products := []map[string]any{
		{"name": "Smartphone", "count": 10, "CategoryID": 1},
		{"name": "Book", "count": 5, "CategoryID": 1},
	}
	for _, product := range products {
		productJSON, _ := json.Marshal(product)
		resp, err := http.Post(cfg.Server.URL+"/api/products", "application/json", bytes.NewReader(productJSON))
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected status 201, got: %v", resp.StatusCode)
		}
	}

	// === 3. Создаём пользователя ===
	user := map[string]string{
		"login":    "testuser",
		"password": "password123",
	}
	userBody, _ := json.Marshal(user)
	resp, err = http.Post(cfg.Server.URL+"/api/users", "application/json", bytes.NewReader(userBody))
	if err != nil {
		t.Fatalf("failed to send user request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	// === 4. Создаём заказ ===
	order := map[string]any{
		"status":  "New",
		"user_id": 1,
		"items": []map[string]any{
			{"productID": 1, "quantity": 2},
			{"productID": 2, "quantity": 1},
		},
	}
	orderBody, _ := json.Marshal(order)
	resp, err = http.Post(cfg.Server.URL+"/api/orders", "application/json", bytes.NewReader(orderBody))
	if err != nil {
		t.Fatalf("failed to send order request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	// === 5. Проверяем наличие заказов и связей в БД ===
	var ordersDB []models.Order
	result := cfg.DB.Raw("SELECT * FROM orders").Scan(&ordersDB)
	if result.Error != nil {
		t.Fatalf("failed to query orders: %v", result.Error)
	}

	if len(ordersDB) != 1 {
		t.Fatalf("expected 1 order, got %d", len(ordersDB))
	}

	var orderProductsDB []models.OrderProduct
	result = cfg.DB.Raw("SELECT * FROM order_products").Scan(&orderProductsDB)
	if result.Error != nil {
		t.Fatalf("failed to query order_products: %v", result.Error)
	}

	if len(orderProductsDB) != 2 {
		t.Fatalf("expected 2 order_products, got %d", len(orderProductsDB))
	}

	t.Log("orders test passed")
}
