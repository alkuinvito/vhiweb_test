package middlewares

import (
	"log"
	"net/http"
	"vhiweb_test/app/users"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	userService *users.UserService
}

type IAuthMiddleware interface {
	Authenticated(c *fiber.Ctx) error
}

func NewAuthMiddleware(userService *users.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService}
}

func (am *AuthMiddleware) Authenticated(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	_, err := am.userService.VerifyToken(authHeader[7:])
	if err != nil {
		log.Fatal(err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Next()
}
