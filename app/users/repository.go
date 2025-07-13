package users

import "gorm.io/gorm"

type UserRepository struct{}

type IUserRepository interface {
	create(db *gorm.DB, user UserModel) (UserModel, error)
	delete(db *gorm.DB, id string) error
	findByEmail(db *gorm.DB, id string) (UserModel, error)
	findById(db *gorm.DB, id string) (UserModel, error)
	get(db *gorm.DB) ([]UserModel, error)
	update(db *gorm.DB, id string, user UserModel) (UserModel, error)
}

func (ur *UserRepository) create(db *gorm.DB, user UserModel) (UserModel, error) {
	err := db.Create(&user).Error

	return user, err
}

func (ur *UserRepository) findByEmail(db *gorm.DB, email string) (UserModel, error) {
	var user UserModel
	err := db.First(&user, "email = ?", email).Error

	return user, err
}

func (ur *UserRepository) findById(db *gorm.DB, id string) (UserModel, error) {
	var user UserModel
	err := db.First(&user, "id = ?", id).Error

	return user, err
}

func (ur *UserRepository) get(db *gorm.DB) ([]UserModel, error) {
	var users []UserModel
	err := db.Find(&users).Error

	return users, err
}
