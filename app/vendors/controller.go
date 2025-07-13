package vendors

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type VendorController struct {
	vendorService *VendorService
}

type IVendorController interface {
	DeleteVendor(c *fiber.Ctx) error
	GetVendorById(c *fiber.Ctx) error
	GetVendorByUserId(c *fiber.Ctx) error
	GetVendors(c *fiber.Ctx) error
	RegisterAsVendor(c *fiber.Ctx) error
	UpdateVendor(c *fiber.Ctx) error
}

func NewVendorController(vendorService *VendorService) *VendorController {
	return &VendorController{vendorService}
}

func (vc *VendorController) DeleteVendor(c *fiber.Ctx) error {
	vendorId := c.Params("id")
	userId := c.Locals("userId").(string)

	err := vc.vendorService.DeleteVendor(vendorId, userId)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "cannot delete vendor"})
		case "not found":
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "vendor not found"})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"message": "vendor deleted successfully"})
}

func (vc *VendorController) GetVendorById(c *fiber.Ctx) error {
	vendorId := c.Params("id")

	vendor, err := vc.vendorService.GetVendorById(vendorId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "vendor not found"})
	}

	return c.JSON(fiber.Map{"vendor": vendor})
}

func (vc *VendorController) GetVendorByUserId(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	vendor, err := vc.vendorService.GetVendorByUserId(userId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "vendor not found"})
	}

	return c.JSON(fiber.Map{"vendor": vendor})
}

func (vc *VendorController) GetVendors(c *fiber.Ctx) error {

	vendors, err := vc.vendorService.GetVendors()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "unknown error"})
	}

	return c.JSON(fiber.Map{"vendors": vendors})
}

func (vc *VendorController) RegisterAsVendor(c *fiber.Ctx) error {
	id := c.Locals("userId").(string)

	var req RegisterVendorRequest
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

	err = vc.vendorService.RegisterAsVendor(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "vendor created successfully"})
}

func (vc *VendorController) UpdateVendor(c *fiber.Ctx) error {
	vendorId := c.Params("id")
	userId := c.Locals("userId").(string)

	var req UpdateVendorRequest
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

	err = vc.vendorService.UpdateVendor(vendorId, userId, req)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "cannot update vendor"})
		case "not found":
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "vendor not found"})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"message": "vendor updated successfully"})
}
