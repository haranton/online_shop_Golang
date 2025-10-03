package repo

import "onlineShop/internal/models"

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

func (repo *Repository) GetProducts() (*[]models.Product, error) {

	var products []models.Product
	if err := repo.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (repo *Repository) UpdateProduct(Product *models.Product) (*models.Product, error) {

	if err := repo.db.Save(Product).Error; err != nil {
		return nil, err
	}
	return Product, nil
}

func (repo *Repository) DeleteProduct(id uint) error {

	var Product models.Product
	if err := repo.db.First(&Product, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&Product).Error; err != nil {
		return err
	}
	return nil
}
