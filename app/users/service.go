package users

import (
	"errors"
	"fmt"
	"vhiweb_test/lib/adapters"
	"vhiweb_test/lib/utils"

	"github.com/gofiber/fiber/v2/log"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type UserService struct {
	db             *gorm.DB
	userRepository *UserRepository
}

type IUserService interface {
	Delete(id string) error
	FindById(id string) (UserModel, error)
	Get() ([]UserModel, error)
	Login(credential UserLoginSchema) (string, error)
	Register(user UserRegisterSchema) (UserModel, error)
	Update(id string, user UserModel) (UserModel, error)
}

func (us *UserService) Get() ([]UserModel, error) {
	return us.userRepository.get(us.db)
}

func (us *UserService) Login(credential UserLoginSchema) (string, error) {
	user, err := us.userRepository.findByEmail(us.db, credential.Email)
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("email and/or password incorrect")
		}

		return "", errors.New("db error")
	}

	if !utils.CheckHash(credential.Password, user.Password) {
		return "", errors.New("email and/or password incorrect")
	}

	return "token_here", nil
}

func (us *UserService) Register(input UserRegisterSchema) (UserModel, error) {
	hashedPassword, err := utils.CreateHash(input.Password)
	if err != nil {
		return UserModel{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := UserModel{
		ID:       ulid.Make().String(),
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		DOB:      input.DOB.String(),
		Password: hashedPassword,
		Role:     "user",
	}

	tx := us.db.Begin()
	defer adapters.CommitOrRollback(tx)

	return us.userRepository.create(tx, user)
}
