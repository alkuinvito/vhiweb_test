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
