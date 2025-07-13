package users

import "time"

type UserRegisterSchema struct {
	Name     string     `json:"name" validate:"required,min=3,max=32"`
	Email    string     `json:"email" validate:"required,email"`
	Phone    string     `json:"phone" validate:"required"`
	DOB      *time.Time `json:"dob" validate:"required"`
	Password string     `json:"password" validate:"required,min=8,max=64"`
}

type UserLoginSchema struct {
	Email    string `json:"email" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type GetUserSchema struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type GetUserProfileSchema struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Phone string     `json:"phone"`
	DOB   *time.Time `json:"dob"`
	Role  string     `json:"role"`
}

type UpdateUserSchema struct {
	Name     string     `json:"name" validate:"omitempty,min=3,max=32"`
	Email    string     `json:"email" validate:"omitempty,email"`
	Phone    string     `json:"phone" validate:"omitempty,numeric"`
	DOB      *time.Time `json:"dob"`
	Password string     `json:"password" validate:"omitempty,min=8,max=64"`
}

type RequestTokenSchema struct {
	AccessToken string `json:"access_token"`
	ExpiredIn   int    `json:"expired_in"`
}
