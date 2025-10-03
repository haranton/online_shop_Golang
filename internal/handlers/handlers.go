package handlers

import (
	"net/http"
	"onlineShop/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// Users
	mux.HandleFunc("/api/users", h.UsersHandler) // GET, POST
	mux.HandleFunc("/api/users/", h.UserHandler) // GET, PUT, DELETE

	// Products
	mux.HandleFunc("/api/products", h.ProductsHandler) // GET, POST
	mux.HandleFunc("/api/products/", h.ProductHandler) // GET, PUT, DELETE

	// Categories
	mux.HandleFunc("/api/categories", h.CategoriesHandler)
	mux.HandleFunc("/api/categories/", h.CategoryHandler)

	// Orders
	mux.HandleFunc("/api/orders", h.OrdersHandler) // GET, POST
	mux.HandleFunc("/api/orders/", h.OrderHandler) // GET, PUT, DELETE
}
