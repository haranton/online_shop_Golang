package service

import (
	"errors"
	"fmt"
	"onlineShop/internal/dto"
	"onlineShop/internal/models"
	"onlineShop/internal/repo"
)

type ProductService interface {
	CreateProduct(product *dto.ProductRequest) (*models.Product, error)
	GetProduct(id uint) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	repo *repo.Repository
}

func NewProductService(r *repo.Repository) ProductService {
	return &productService{repo: r}
}

func (s *productService) CreateProduct(product *dto.ProductRequest) (*models.Product, error) {

	if product.Count < 0 {
		return nil, errors.New("count cannot be negative")
	}

	if product.Name == "" {
		return nil, fmt.Errorf("category name cannot be empty")
	}

	newProduct := models.Product{
		Name:       product.Name,
		CategoryID: product.CategoryID,
		Count:      product.Count,
	}

	return s.repo.CreateProduct(&newProduct)
}

func (s *productService) GetProduct(id uint) (*models.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *productService) GetProducts() ([]models.Product, error) {
	products, err := s.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return *products, nil
}

func (s *productService) UpdateProduct(product *models.Product) (*models.Product, error) {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}
