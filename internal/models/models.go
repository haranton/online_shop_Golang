package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	Orders []Order `gorm:"foreignKey:UserID"`
}

type Order struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Status string `gorm:"size:50;not null"`

	UserID uint `gorm:"not null"`
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	Products []OrderProduct `gorm:"foreignKey:OrderID"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:255;unique;not null"`

	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	ID         uint     `gorm:"primaryKey;autoIncrement"`
	Name       string   `gorm:"not null"`
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"constraint:OnDelete:CASCADE;"`

	OrderProducts []OrderProduct `gorm:"foreignKey:ProductID"`
}

type OrderProduct struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  int  `gorm:"not null"`

	Order   Order   `gorm:"constraint:OnDelete:CASCADE;"`
	Product Product `gorm:"constraint:OnDelete:CASCADE;"`
}
