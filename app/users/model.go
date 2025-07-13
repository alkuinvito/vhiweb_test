package users

import (
	"time"
	"vhiweb_test/app/vendors"
)

type UserModel struct {
	ID        string               `gorm:"primaryKey"`
	Name      string               `gorm:"not null"`
	DOB       time.Time            `gorm:"not null"`
	Email     string               `gorm:"not null;uniqueIndex"`
	Password  string               `gorm:"not null"`
	Phone     string               `gorm:"not null;unique"`
	Role      string               `gorm:"not null"`
	Vendor    *vendors.VendorModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
