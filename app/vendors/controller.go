package vendors

import "gorm.io/gorm"

type VendorController struct {
	vendorService *VendorService
}

type IVendorController interface{}

func NewVendorController(db *gorm.DB) *VendorController {
	vendorRepository := &VendorRepository{}
	vendorService := &VendorService{db, vendorRepository}
	return &VendorController{vendorService}
}
