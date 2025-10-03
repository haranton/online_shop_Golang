package repo

import "onlineShop/internal/models"

func (repo *Repository) CreateUser(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
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

	if err := repo.db.Save(User).Error; err != nil {
		return nil, err
	}
	return User, nil
}

func (repo *Repository) DeleteUser(id uint) error {

	var User models.User
	if err := repo.db.First(&User, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&User).Error; err != nil {
		return err
	}
	return nil
}
