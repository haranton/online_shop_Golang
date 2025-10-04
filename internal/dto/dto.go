package dto

import "onlineShop/internal/models"

type Items struct {
	ProductID uint
	Quantity  int
}

// Запрос на обновление заказа (например, только статус)
type OrderUpdateRequest struct {
	Status models.StatusOrder `json:"status"`
}

type OrderCreateRequest struct {
	UserID uint    `json:"user_id"`
	Items  []Items `json:"items"`
}
