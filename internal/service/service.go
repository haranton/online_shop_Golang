package service

import "onlineShop/internal/repo"

// Service — фасад для всех сервисов
type Service struct {
	User     UserService
	Product  ProductService
	Category CategoryService
	Order    OrderService
}

// NewService — конструктор, который собирает все сервисы
func NewService(r *repo.Repository) *Service {
	return &Service{
		User:     NewUserService(r),
		Product:  NewProductService(r),
		Category: NewCategoryService(r),
		Order:    NewOrderService(r),
	}
}
