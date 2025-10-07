package dto

import "onlineShop/internal/models"

// --- USERS ---

type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Login string `json:"login"`
}

// --- CATEGORIES ---

type CategoryRequest struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// --- PRODUCTS ---

type ProductRequest struct {
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
	Count      int    `json:"count"`
}

type ProductResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
	Count      int    `json:"count"`
}

// --- ORDERS ---

// Элемент заказа (товар и количество)
type Items struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

// Создание нового заказа
type OrderCreateRequest struct {
	UserID uint    `json:"user_id"`
	Items  []Items `json:"items"`
}

// Обновление существующего заказа
type OrderUpdateRequest struct {
	Status models.StatusOrder `json:"status"`
}

// Ответ при получении заказа
type OrderResponse struct {
	ID     uint               `json:"id"`
	Status models.StatusOrder `json:"status"`
	UserID uint               `json:"user_id"`
	Items  []OrderItemDetail  `json:"items,omitempty"`
}

// Детали товара в заказе (для ответа)
type OrderItemDetail struct {
	ProductID uint   `json:"product_id"`
	Name      string `json:"name,omitempty"`
	Quantity  int    `json:"quantity"`
}
