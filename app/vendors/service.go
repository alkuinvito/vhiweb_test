package vendors

import (
	"errors"
	"vhiweb_test/lib/adapters"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type VendorService struct {
	db               *gorm.DB
	vendorRepository *VendorRepository
}

type IVendorService interface {
	DeleteVendor(id, userId string) error
	GetVendorById(id string) (GetVendorResponse, error)
	GetVendorByUserId(userId string) (GetVendorResponse, error)
	GetVendors() ([]GetVendorResponse, error)
	RegisterAsVendor(userId string, input RegisterVendorRequest) error
	UpdateVendor(id, userId string, input UpdateVendorRequest) error
}

func NewVendorService(db *gorm.DB, vendorRepository *VendorRepository) *VendorService {
	return &VendorService{db, vendorRepository}
}

func (vs *VendorService) DeleteVendor(id, userId string) error {
	tx := vs.db.Begin()
	defer adapters.CommitOrRollback(tx)

	vendor, err := vs.vendorRepository.findByUserId(tx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found")
		}

		return err
	}

	if vendor.UserID != userId {
		return errors.New("unauthorized")
	}

	return vs.vendorRepository.delete(tx, id)
}

func (vs *VendorService) GetVendorById(id string) (GetVendorResponse, error) {
	var result GetVendorResponse

	tx := vs.db.Begin()
	defer adapters.CommitOrRollback(tx)

	vendor, err := vs.vendorRepository.findById(tx, id)
	if err != nil {
		return result, err
	}

	result.ID = vendor.ID
	result.Name = vendor.Name
	result.UserID = vendor.UserID

	return result, nil
}

func (vs *VendorService) GetVendorByUserId(userId string) (GetVendorResponse, error) {
	var result GetVendorResponse

	tx := vs.db.Begin()
	defer adapters.CommitOrRollback(tx)

	vendor, err := vs.vendorRepository.findByUserId(tx, userId)
	if err != nil {
		return result, err
	}

	result.ID = vendor.ID
	result.Name = vendor.Name
	result.UserID = vendor.UserID

	return result, nil
}

func (vs *VendorService) GetVendors() ([]GetVendorResponse, error) {
	var result []GetVendorResponse

	tx := vs.db.Begin()
	defer adapters.CommitOrRollback(tx)

	vendors, err := vs.vendorRepository.get(tx)
	if err != nil {
		return result, err
	}

	for _, vendor := range vendors {
		result = append(result, GetVendorResponse{
			ID:     vendor.ID,
			Name:   vendor.Name,
			UserID: vendor.UserID,
		})
	}

	return result, nil
}

func (vs *VendorService) RegisterAsVendor(userId string, input RegisterVendorRequest) error {
	_, err := vs.vendorRepository.findByUserId(vs.db, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			vendor := VendorModel{
				ID:     ulid.Make().String(),
				Name:   input.Name,
				UserID: userId,
			}

			_, err := vs.vendorRepository.create(vs.db, vendor)
			if err != nil {
				return err
			}

			return nil
		}

		return errors.New("db error")
	}

	return errors.New("user is already registered as vendor")
}

func (vs *VendorService) UpdateVendor(id, userId string, input UpdateVendorRequest) error {
	tx := vs.db.Begin()
	defer adapters.CommitOrRollback(tx)

	vendor, err := vs.vendorRepository.findByUserId(tx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found")
		}

		return err
	}

	if vendor.UserID != userId {
		return errors.New("unauthorized")
	}

	updated := VendorModel{ID: id, Name: input.Name}

	err = vs.vendorRepository.update(tx, updated)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found")
		}

		return err
	}

	return nil
}
