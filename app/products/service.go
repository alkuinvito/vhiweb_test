package products

import (
	"errors"
	"vhiweb_test/lib/adapters"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type ProductService struct {
	db                *gorm.DB
	productRepository *ProductRepository
}

type IProductService interface {
	CreateProduct(vendorId string, input CreateProductRequest) error
	DeleteProduct(id, vendorId string) error
	GetProductById(id string) (GetProductResponse, error)
	GetProductsByUserId(vendorId string) ([]GetProductResponse, error)
	GetProducts() ([]GetProductResponse, error)
	UpdateProduct(id, vendorId string, input UpdateProductRequest) error
}

func NewProductService(db *gorm.DB, productRepository *ProductRepository) *ProductService {
	return &ProductService{db, productRepository}
}

func (ps *ProductService) CreateProduct(vendorId string, input CreateProductRequest) error {
	tx := ps.db.Begin()
	defer adapters.CommitOrRollback(tx)

	product := ProductModel{
		ID:          ulid.Make().String(),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		VendorID:    vendorId,
	}

	_, err := ps.productRepository.create(tx, product)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) DeleteProduct(id, vendorId string) error {
	tx := ps.db.Begin()
	defer adapters.CommitOrRollback(tx)

	product, err := ps.productRepository.findById(tx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found")
		}

		return err
	}

	if product.VendorID != vendorId {
		return errors.New("unauthorized")
	}

	return ps.productRepository.delete(tx, id)
}

func (ps *ProductService) GetProductById(id string) (GetProductResponse, error) {
	var result GetProductResponse

	tx := ps.db.Begin()
	defer adapters.CommitOrRollback(tx)

	product, err := ps.productRepository.findById(tx, id)
	if err != nil {
		return result, err
	}

	result.ID = product.ID
	result.Name = product.Name
	result.Description = product.Description
	result.Price = product.Price
	result.VendorID = product.VendorID

	return result, nil
}

func (ps *ProductService) GetProductsByUserId(vendorId string) ([]GetProductResponse, error) {
	var result []GetProductResponse

	tx := ps.db.Begin()
	defer adapters.CommitOrRollback(tx)

	products, err := ps.productRepository.findByUserId(tx, vendorId)
	if err != nil {
		return result, err
	}

	for _, product := range products {
		result = append(result, GetProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			VendorID:    product.VendorID,
		})
	}

	return result, nil
}

func (ps *ProductService) GetProducts() ([]GetProductResponse, error) {
	var result []GetProductResponse

	tx := ps.db.Begin()
	defer adapters.CommitOrRollback(tx)

	products, err := ps.productRepository.get(tx)
	if err != nil {
		return result, err
	}

	for _, product := range products {
		result = append(result, GetProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			VendorID:    product.VendorID,
		})
	}

	return result, nil
}

func (ps *ProductService) UpdateProduct(id, vendorId string, input UpdateProductRequest) error {
	tx := ps.db.Begin()
	defer adapters.CommitOrRollback(tx)

	product, err := ps.productRepository.findById(tx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found")
		}

		return err
	}

	if product.VendorID != vendorId {
		return errors.New("unauthorized")
	}

	updated := ProductModel{ID: id, Name: input.Name}

	err = ps.productRepository.update(tx, updated)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found")
		}

		return err
	}

	return nil
}
