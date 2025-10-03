package repo

import "onlineShop/internal/models"

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

func (repo *Repository) GetCategories() (*[]models.Category, error) {

	var categories []models.Category
	if err := repo.db.Find(&categories).Error; err != nil {
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

func (repo *Repository) DeleteCategory(id uint) error {

	var category models.Category
	if err := repo.db.First(&category, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}
