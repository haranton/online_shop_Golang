package service

import (
	"fmt"
	"onlineShop/internal/models"
	"onlineShop/internal/repo"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(id uint) (*models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo *repo.Repository
}

func NewUserService(r *repo.Repository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	if user.Login == "" {
		return nil, fmt.Errorf("email is required")
	}
	return s.repo.CreateUser(user)
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *userService) GetUsers() ([]models.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return *users, nil
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
