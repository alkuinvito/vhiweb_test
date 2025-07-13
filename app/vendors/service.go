package vendors

import "gorm.io/gorm"

type VendorService struct {
	db               *gorm.DB
	vendorRepository *VendorRepository
}

type IVendorService interface{}
