package products

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService *ProductService
}

type IProductController interface {
	CreateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	GetProductsByUserId(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
}

func NewProductController(productService *ProductService) *ProductController {
	return &ProductController{productService}
}

func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
	vendorId := c.Locals("vendorId").(string)

	var req CreateProductRequest
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

	err = pc.productService.CreateProduct(vendorId, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "product created successfully"})
}

func (pc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	vendorId := c.Locals("vendorId").(string)

	err := pc.productService.DeleteProduct(productId, vendorId)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "cannot delete product"})
		case "not found":
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"message": "product deleted successfully"})
}

func (pc *ProductController) GetProductById(c *fiber.Ctx) error {
	productId := c.Params("id")

	product, err := pc.productService.GetProductById(productId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}

	return c.JSON(fiber.Map{"product": product})
}

func (pc *ProductController) GetProductsByUserId(c *fiber.Ctx) error {
	vendorId := c.Locals("vendorId").(string)

	products, err := pc.productService.GetProductsByUserId(vendorId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}

	return c.JSON(fiber.Map{"product": products})
}

func (pc *ProductController) GetProducts(c *fiber.Ctx) error {
	products, err := pc.productService.GetProducts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "unknown error"})
	}

	return c.JSON(fiber.Map{"products": products})
}

func (pc *ProductController) UpdateProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	vendorId := c.Locals("vendorId").(string)

	var req UpdateProductRequest
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

	err = pc.productService.UpdateProduct(productId, vendorId, req)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "cannot update product"})
		case "not found":
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"message": "product updated successfully"})
}
