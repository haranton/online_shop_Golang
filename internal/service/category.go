package service

import (
	"fmt"
	"onlineShop/internal/models"
	"onlineShop/internal/repo"
)

type CategoryService interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	GetCategory(id uint) (*models.Category, error)
	GetCategories() ([]models.Category, error)
	UpdateCategory(category *models.Category) (*models.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo *repo.Repository
}

func NewCategoryService(r *repo.Repository) CategoryService {
	return &categoryService{repo: r}
}

func (s *categoryService) CreateCategory(category *models.Category) (*models.Category, error) {

	if category.Name == "" {
		return nil, fmt.Errorf("category name cannot be empty")
	}
	return s.repo.CreateCategories(category)
}

func (s *categoryService) GetCategory(id uint) (*models.Category, error) {
	return s.repo.GetCategory(id)
}

func (s *categoryService) GetCategories() ([]models.Category, error) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, err
	}
	return *categories, nil
}

func (s *categoryService) UpdateCategory(category *models.Category) (*models.Category, error) {
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.DeleteCategory(id)
}
