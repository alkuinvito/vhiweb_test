package middlewares

import (
	"net/http"
	"vhiweb_test/app/users"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	userService *users.UserService
}

type IAuthMiddleware interface {
	Authenticated(c *fiber.Ctx) error
	AuthorizedSubject(c *fiber.Ctx) error
}

func NewAuthMiddleware(userService *users.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService}
}

func (am *AuthMiddleware) Authenticated(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token, err := am.userService.VerifyToken(authHeader[7:])
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Unknown error"})
	}

	c.Locals("userId", sub)

	return c.Next()
}

func (am *AuthMiddleware) AuthorizedSubject(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token, err := am.userService.VerifyToken(authHeader[7:])
	if err != nil {

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Unknown error"})
	}

	if sub != id {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	c.Locals("userId", sub)

	return c.Next()
}
