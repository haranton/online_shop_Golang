package repo

import (
	"onlineShop/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewReposytory(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

//Categories

func (repo *Repository) CreateCategories(category *models.Category) (*models.Category, error) {
	if err := repo.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *Repository) GetCategory(id uint) (*models.Category, error) {

	var category models.Category
	if err := repo.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *Repository) GetCategories(id uint) (*[]models.Category, error) {

	var categories []models.Category
	if err := repo.db.Find(&categories, id).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (repo *Repository) UpdateCategory(category *models.Category) (*models.Category, error) {

	if err := repo.db.Save(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *Repository) DeleteCategory(category *models.Category) error {

	if err := repo.db.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

// Products
func (repo *Repository) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := repo.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *Repository) GetProduct(id uint) (*models.Product, error) {

	var product models.Product
	if err := repo.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *Repository) GetProducts(id uint) (*[]models.Product, error) {

	var products []models.Product
	if err := repo.db.Find(&products, id).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (repo *Repository) UpdateProduct(Product *models.Product) (*models.Product, error) {

	if err := repo.db.Save(&Product).Error; err != nil {
		return nil, err
	}
	return Product, nil
}

func (repo *Repository) DeleteProduct(Product *models.Product) error {

	if err := repo.db.Delete(&Product).Error; err != nil {
		return err
	}
	return nil
}

// Users
func (repo *Repository) CreateUser(category *models.User) (*models.User, error) {
	if err := repo.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *Repository) GetUser(id uint) (*models.User, error) {

	var category models.User
	if err := repo.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *Repository) GetUsers(id uint) (*[]models.User, error) {

	var users []models.User
	if err := repo.db.Find(&users, id).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (repo *Repository) UpdateUser(User *models.User) (*models.User, error) {

	if err := repo.db.Save(&User).Error; err != nil {
		return nil, err
	}
	return User, nil
}

func (repo *Repository) DeleteUser(User *models.User) error {

	if err := repo.db.Delete(&User).Error; err != nil {
		return err
	}
	return nil
}

//
