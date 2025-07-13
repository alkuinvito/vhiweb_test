package vendors

import "time"

type VendorModel struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	UserID    string `gorm:"not null;uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
