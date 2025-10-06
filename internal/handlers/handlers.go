package handlers

import (
	"net/http"
	"onlineShop/internal/service"

	_ "onlineShop/docs"

	httpSwagger "github.com/swaggo/http-swagger"
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

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Users
	mux.HandleFunc("GET /api/users", h.UsersGetHandler)
	mux.HandleFunc("POST /api/users", h.UsersPostHandler)

	mux.HandleFunc("GET /api/users/{id}", h.UserGetHandler)
	mux.HandleFunc("PUT /api/users/{id}", h.UserPutHandler)
	mux.HandleFunc("DELETE /api/users/{id}", h.UserDeleteHandler)

	// Products
	mux.HandleFunc("GET /api/products", h.ProductsGetHandler)
	mux.HandleFunc("POST /api/products", h.ProductsPostHandler)

	mux.HandleFunc("GET /api/products/{id}", h.ProductGetHandler)
	mux.HandleFunc("PUT /api/products/{id}", h.ProductPutHandler)
	mux.HandleFunc("DELETE /api/products/{id}", h.ProductDeleteHandler)

	// Categories
	mux.HandleFunc("GET /api/categories", h.CategoriesGetHandler)
	mux.HandleFunc("POST /api/categories", h.CategoriesPostHandler)

	mux.HandleFunc("GET /api/categories/{id}", h.CategoryGetHandler)
	mux.HandleFunc("PUT /api/categories/{id}", h.CategoryPutHandler)
	mux.HandleFunc("DELETE /api/categories/{id}", h.CategoryDeleteHandler)

	// Orders
	mux.HandleFunc("POST /api/orders", h.OrdersPostHandler)

	mux.HandleFunc("GET /api/orders/{id}", h.OrderGetHandler)
	mux.HandleFunc("PUT /api/orders/{id}", h.OrderPutHandler)
	mux.HandleFunc("DELETE /api/orders/{id}", h.OrderDeleteHandler)

	mux.HandleFunc("GET /api/users/{id}/orders", h.GetOrdersByUserIDHandler)
	mux.HandleFunc("DELETE /api/orders/{orderId}/products/{productId}", h.DeleteProductFromOrderHandler)
}
