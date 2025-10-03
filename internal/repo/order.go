package repo

import (
	"fmt"
	"onlineShop/internal/dto"
	"onlineShop/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *Repository) CreateOrder(userID uint, items []dto.Items) (*models.Order, error) {

	var order models.Order

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		order = models.Order{
			Status: string(models.StatusNew),
			UserID: user.ID,
		}

		if err := tx.Create(&order).Error; err != nil {
			return fmt.Errorf("failed to create order: %w", err)
		}

		for _, item := range items {

			var product models.Product

			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, item.ProductID).Error; err != nil {
				return fmt.Errorf("product %d not found", item.ProductID)
			}

			if product.Count < item.Quantity {
				return fmt.Errorf("not enough stock for product %d", product.ID)
			}

			if err := tx.Model(&product).
				Where("count >= ?", item.Quantity).
				Update("count", gorm.Expr("count - ?", item.Quantity)).Error; err != nil {
				return err
			}

			orderProduct := models.OrderProduct{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  item.Quantity,
			}

			if err := tx.Create(&orderProduct).Error; err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (repo *Repository) GetOrder(id uint) (*models.Order, error) {

	var order models.Order
	if err := repo.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *Repository) GetOrders() (*[]models.Order, error) {

	var orders []models.Order
	if err := repo.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

func (repo *Repository) UpdateStatusOrder(id uint, status models.StatusOrder) (*models.Order, error) {

	var order models.Order
	if err := repo.db.First(&order, id).Error; err != nil {
		return nil, err
	}

	order.Status = string(status)
	if err := repo.db.Save(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *Repository) DeleteOrder(id uint) error {
	var order models.Order
	if err := repo.db.First(&order, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetOrdersByUserID(userID uint) (*[]models.Order, error) {

	var orders []models.Order
	if err := repo.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return &orders, nil
}

func (repo *Repository) DeleteProductFromOrder(orderID uint, productID uint) error {
	var orderProduct models.OrderProduct
	if err := repo.db.Where("order_id = ? AND product_id = ?", orderID, productID).First(&orderProduct).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&orderProduct).Error; err != nil {
		return err
	}
	return nil
}
