package users

import (
	"errors"
	"fmt"
	"os"
	"time"
	"vhiweb_test/lib/adapters"
	"vhiweb_test/lib/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

var EXPIRED_IN = time.Hour * 24 // expired in a day

type UserService struct {
	db             *gorm.DB
	userRepository *UserRepository
}

type IUserService interface {
	DeleteUser(id string) error
	generateToken(user UserModel) (string, error)
	GetUserById(id string) (GetUserSchema, error)
	GetUserProfile(id string) (GetUserProfileSchema, error)
	GetUsers() ([]GetUserSchema, error)
	Login(credential UserLoginSchema) (string, error)
	Register(user UserRegisterSchema) (GetUserProfileSchema, error)
	UpdateUser(id string, input UpdateUserSchema) error
	VerifyToken(tokenString string) (any, error)
}

func (us *UserService) DeleteUser(id string) error {
	return us.userRepository.delete(us.db, id)
}

func (us *UserService) generateToken(user UserModel) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(EXPIRED_IN).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("failed to generate JWT")
	}

	return signed, nil
}

func (us *UserService) GetUserById(id string) (GetUserSchema, error) {
	var user GetUserSchema

	result, err := us.userRepository.findById(us.db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}

		return user, errors.New("db error")
	}

	user.ID = result.ID
	user.Name = result.Name
	user.Phone = result.Phone
	user.Role = result.Role

	return user, nil
}

func (us *UserService) GetUserProfile(id string) (GetUserProfileSchema, error) {
	var user GetUserProfileSchema

	result, err := us.userRepository.findById(us.db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}

		return user, errors.New("db error")
	}

	timeLayout := "2006-01-02 15:04:05 -0700 MST"
	parsedDOB, err := time.Parse(timeLayout, result.DOB)
	if err != nil {
		return user, err
	}

	user.ID = result.ID
	user.Name = result.Name
	user.Email = result.Email
	user.Phone = result.Phone
	user.DOB = &parsedDOB
	user.Role = result.Role

	return user, nil
}

func (us *UserService) GetUsers() ([]GetUserSchema, error) {
	var users []GetUserSchema

	results, err := us.userRepository.get(us.db)
	if err != nil {
		return users, err
	}

	for _, user := range results {
		users = append(users, GetUserSchema{
			ID:    user.ID,
			Name:  user.Name,
			Phone: user.Phone,
			Role:  user.Role,
		})
	}

	return users, nil
}

func (us *UserService) Login(credential UserLoginSchema) (RequestTokenSchema, error) {
	var token RequestTokenSchema

	user, err := us.userRepository.findByEmail(us.db, credential.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return token, errors.New("email and/or password incorrect")
		}

		return token, errors.New("db error")
	}

	if !utils.CheckHash(credential.Password, user.Password) {
		return token, errors.New("email and/or password incorrect")
	}

	accessToken, err := us.generateToken(user)
	if err != nil {
		return token, err
	}

	token.AccessToken = accessToken
	token.ExpiredIn = int(EXPIRED_IN.Seconds())

	return token, nil
}

func (us *UserService) Register(input UserRegisterSchema) (GetUserProfileSchema, error) {
	var updated GetUserProfileSchema
	hashedPassword, err := utils.CreateHash(input.Password)
	if err != nil {
		return updated, fmt.Errorf("failed to hash password: %w", err)
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

	result, err := us.userRepository.create(tx, user)
	if err != nil {
		return updated, err
	}

	timeLayout := "2006-01-02 15:04:05 -0700 MST"
	parsedDOB, err := time.Parse(timeLayout, result.DOB)
	if err != nil {
		return updated, err
	}

	updated.ID = result.ID
	updated.Name = result.Name
	updated.Email = result.Email
	updated.Phone = result.Phone
	updated.DOB = &parsedDOB
	updated.Role = result.Role

	return updated, nil
}

func (us *UserService) UpdateUser(id string, input UpdateUserSchema) error {
	var err error
	var user UserModel

	if input.Password != "" {
		input.Password, err = utils.CreateHash(input.Password)
		if err != nil {
			return err
		}
	}

	if input.DOB != nil {
		user.DOB = input.DOB.String()
	}

	user = UserModel{
		ID:       id,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	}

	tx := us.db.Begin()
	defer adapters.CommitOrRollback(tx)

	return us.userRepository.update(tx, user)
}

func (us *UserService) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	return token, nil
}
