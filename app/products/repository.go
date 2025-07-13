package products

import (
	"gorm.io/gorm"
)

type ProductRepository struct{}

type IProductRepository interface {
	create(db *gorm.DB, product ProductModel) (ProductModel, error)
	delete(db *gorm.DB, id string) error
	findById(db *gorm.DB, id string) (ProductModel, error)
	findByUserId(db *gorm.DB, userId string) ([]ProductModel, error)
	get(db *gorm.DB) ([]ProductModel, error)
	update(db *gorm.DB, product ProductModel) error
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (pr *ProductRepository) create(db *gorm.DB, product ProductModel) (ProductModel, error) {
	err := db.Create(&product).Error

	return product, err
}

func (pr *ProductRepository) delete(db *gorm.DB, id string) error {
	err := db.Delete(&ProductModel{}, "id = ?", id).Error

	return err
}

func (pr *ProductRepository) findById(db *gorm.DB, id string) (ProductModel, error) {
	var product ProductModel
	err := db.First(&product, "id = ?", id).Error

	return product, err
}

func (pr *ProductRepository) findByUserId(db *gorm.DB, userId string) ([]ProductModel, error) {
	var products []ProductModel
	err := db.Find(&products, "user_id = ?", userId).Error

	return products, err
}

func (pr *ProductRepository) get(db *gorm.DB) ([]ProductModel, error) {
	var products []ProductModel
	err := db.Find(&products).Error

	return products, err
}

func (pr *ProductRepository) update(db *gorm.DB, product ProductModel) error {
	return db.Model(&product).Updates(product).Error
}
