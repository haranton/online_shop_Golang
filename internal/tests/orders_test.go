package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"onlineShop/internal/models"
	"testing"
)

func TestOrders(t *testing.T) {
	cfg, err := SetupTestEnv()
	if err != nil {
		t.Fatalf("failed to load env: %v", err)
	}

	t.Log(" Starting orders integration test")

	// === 1. Создаём категорию ===
	category := map[string]string{"name": "Electronics"}
	body, _ := json.Marshal(category)
	resp, err := http.Post(cfg.Server.URL+"/api/categories", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to send category request: %v", err)
	}
	defer resp.Body.Close()

	var categoryResp struct {
		ID uint `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&categoryResp); err != nil {
		t.Fatalf("failed to decode category response: %v", err)
	}
	t.Logf("Category created with ID %d", categoryResp.ID)

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected 201, got %d, body: %s", resp.StatusCode, string(respBody))
	}
	t.Log("Category created")

	// === 2. Создаём продукты ===
	products := []map[string]any{
		{"name": "Smartphone", "count": 10, "category_id": 1},
		{"name": "Book", "count": 5, "category_id": 1},
	}

	for _, product := range products {
		productJSON, _ := json.Marshal(product)
		resp, err := http.Post(cfg.Server.URL+"/api/products", "application/json", bytes.NewReader(productJSON))
		if err != nil {
			t.Fatalf("failed to send product request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			respBody, _ := io.ReadAll(resp.Body)
			t.Fatalf("expected 201, got %d, body: %s", resp.StatusCode, string(respBody))
		}
	}
	t.Log("Products created")

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
		respBody, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected 201, got %d, body: %s", resp.StatusCode, string(respBody))
	}
	t.Log("User created")

	// === 4. Создаём заказ ===
	order := map[string]any{
		"user_id": 1,
		"items": []map[string]any{
			{"product_id": 1, "quantity": 2},
			{"product_id": 2, "quantity": 1},
		},
	}
	orderBody, _ := json.Marshal(order)
	resp, err = http.Post(cfg.Server.URL+"/api/orders", "application/json", bytes.NewReader(orderBody))
	if err != nil {
		t.Fatalf("failed to send order request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected 201, got %d, body: %s", resp.StatusCode, string(respBody))
	}
	t.Log("Order created")

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

	t.Log("Orders and related products verified")
	t.Log("Orders integration test passed successfully")
}
