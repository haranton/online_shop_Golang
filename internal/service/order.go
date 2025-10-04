package service

import (
	"fmt"
	"onlineShop/internal/dto"
	"onlineShop/internal/models"
	"onlineShop/internal/repo"
)

type OrderService interface {
	CreateOrder(userID uint, items []dto.Items) (*models.Order, error)
	GetOrder(id uint) (*models.Order, error)
	GetOrdersByUserID(userID uint) (*[]models.Order, error)
	UpdateStatusOrder(id uint, status models.StatusOrder) (*models.Order, error)
	DeleteOrder(id uint) error
	DeleteProductFromOrder(orderID, productID uint) error
}

type orderService struct {
	repo *repo.Repository
}

func NewOrderService(r *repo.Repository) OrderService {
	return &orderService{repo: r}
}

func (s *orderService) CreateOrder(userID uint, items []dto.Items) (*models.Order, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("order must contain at least one item")
	}
	return s.repo.CreateOrder(userID, items)
}

func (s *orderService) GetOrder(id uint) (*models.Order, error) {
	return s.repo.GetOrder(id)
}

func (s *orderService) GetOrdersByUserID(userID uint) (*[]models.Order, error) {
	return s.repo.GetOrdersByUserID(userID)
}

func (s *orderService) UpdateStatusOrder(id uint, status models.StatusOrder) (*models.Order, error) {
	return s.repo.UpdateStatusOrder(id, status)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.repo.DeleteOrder(id)
}

func (s *orderService) DeleteProductFromOrder(orderID, productID uint) error {
	return s.repo.DeleteProductFromOrder(orderID, productID)
}
