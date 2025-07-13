package products

import "time"

type ProductModel struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	VendorID    string `gorm:"not null;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
