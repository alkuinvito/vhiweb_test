package vendors

import (
	"time"
	"vhiweb_test/app/products"
)

type VendorModel struct {
	ID        string                   `gorm:"primaryKey"`
	Name      string                   `gorm:"not null"`
	UserID    string                   `gorm:"not null;uniqueIndex"`
	Products  *[]products.ProductModel `gorm:"foreignKey:VendorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
