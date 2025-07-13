package vendors

import (
	"gorm.io/gorm"
)

type VendorRepository struct{}

type IVendorRepository interface {
	create(db *gorm.DB, vendor VendorModel) (VendorModel, error)
	delete(db *gorm.DB, id string) error
	findById(db *gorm.DB, id string) (VendorModel, error)
	findByUserId(db *gorm.DB, userId string) (VendorModel, error)
	get(db *gorm.DB) ([]VendorModel, error)
	update(db *gorm.DB, vendor VendorModel) error
}

func NewVendorRepository() *VendorRepository {
	return &VendorRepository{}
}

func (ur *VendorRepository) create(db *gorm.DB, vendor VendorModel) (VendorModel, error) {
	err := db.Create(&vendor).Error

	return vendor, err
}

func (ur *VendorRepository) delete(db *gorm.DB, id string) error {
	err := db.Delete(&VendorModel{}, "id = ?", id).Error

	return err
}

func (ur *VendorRepository) findById(db *gorm.DB, id string) (VendorModel, error) {
	var vendor VendorModel
	err := db.First(&vendor, "id = ?", id).Error

	return vendor, err
}

func (ur *VendorRepository) findByUserId(db *gorm.DB, userId string) (VendorModel, error) {
	var vendor VendorModel
	err := db.First(&vendor, "user_id = ?", userId).Error

	return vendor, err
}

func (ur *VendorRepository) get(db *gorm.DB) ([]VendorModel, error) {
	var vendors []VendorModel
	err := db.Find(&vendors).Error

	return vendors, err
}

func (ur *VendorRepository) update(db *gorm.DB, vendor VendorModel) error {
	return db.Model(&vendor).Updates(vendor).Error
}
