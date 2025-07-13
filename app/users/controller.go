package users

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	userService *UserService
}

type IUserController interface {
	DeleteUser(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	GetUserProfile(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

func NewUserController(db *gorm.DB) *UserController {
	userRepository := &UserRepository{}
	userService := &UserService{db, userRepository}
	return &UserController{userService}
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := uc.userService.DeleteUser(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "user deleted successfully"})
}

func (uc *UserController) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uc.userService.GetUserById(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"user": user})
}

func (uc *UserController) GetUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uc.userService.GetUserProfile(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"user": user})
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.userService.GetUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"users": users})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var req UserLoginSchema
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
	}

	token, err := uc.userService.Login(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"access_token": token})
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var req UserRegisterSchema
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
	}

	user, err := uc.userService.Register(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"user": user})
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var req UpdateUserSchema
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
	}

	err = uc.userService.UpdateUser(id, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "user updated successfully"})
}
