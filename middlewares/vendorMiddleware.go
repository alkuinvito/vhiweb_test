package middlewares

import (
	"net/http"
	"vhiweb_test/app/users"
	"vhiweb_test/app/vendors"

	"github.com/gofiber/fiber/v2"
)

type VendorMiddleware struct {
	userService   *users.UserService
	vendorService *vendors.VendorService
}

type IVendorMiddleware interface {
	RegisteredVendor(c *fiber.Ctx) error
}

func NewVendorMiddleware(userService *users.UserService, vendorService *vendors.VendorService) *VendorMiddleware {
	return &VendorMiddleware{userService, vendorService}
}

func (vm *VendorMiddleware) RegisteredVendor(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token, err := vm.userService.VerifyToken(authHeader[7:])
	if err != nil {

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Unknown error"})
	}

	vendor, err := vm.vendorService.GetVendorByUserId(sub)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "For vendor only"})
	}

	c.Locals("userId", sub)
	c.Locals("vendorId", vendor.ID)

	return c.Next()
}
