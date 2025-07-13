package routers

import (
	"vhiweb_test/app/products"
	"vhiweb_test/app/users"
	"vhiweb_test/app/vendors"
	"vhiweb_test/lib/adapters"
	"vhiweb_test/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Handle(app *fiber.App) {
	db := adapters.NewDB()

	userRepository := users.NewUserRepository()
	userService := users.NewUserService(db, userRepository)
	userController := users.NewUserController(userService)

	vendorRepository := vendors.NewVendorRepository()
	vendorService := vendors.NewVendorService(db, vendorRepository)
	vendorController := vendors.NewVendorController(vendorService)

	productRepository := products.NewProductRepository()
	productService := products.NewProductService(db, productRepository)
	productController := products.NewProductController(productService)

	authMiddleware := middlewares.NewAuthMiddleware(userService)
	vendorMiddleware := middlewares.NewVendorMiddleware(userService, vendorService)

	app.Use(recover.New())
	app.Use(cors.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	userRouter := v1.Group("/users")
	userRouter.Get("/", userController.GetUsers)
	userRouter.Get("/:id", userController.GetUserById)
	// authenticated user only
	userRouter.Use(authMiddleware.AuthorizedSubject)
	userRouter.Patch("/:id", userController.UpdateUser)
	userRouter.Delete("/:id", userController.DeleteUser)

	authRouter := v1.Group("/auth")
	authRouter.Post("/login", userController.Login)
	authRouter.Post("/register", userController.Register)

	vendorRouter := v1.Group("/vendors")
	vendorRouter.Get("/", vendorController.GetVendors)
	vendorRouter.Get("/:id", vendorController.GetVendorById)
	// authenticated user only
	vendorRouter.Use(authMiddleware.Authenticated)
	vendorRouter.Post("/", vendorController.RegisterAsVendor)
	vendorRouter.Patch("/:id", vendorController.UpdateVendor)
	vendorRouter.Delete("/:id", vendorController.DeleteVendor)

	productRouter := v1.Group("/products")
	productRouter.Get("/", productController.GetProducts)
	productRouter.Get("/:id", productController.GetProductById)
	// registered vendor only
	productRouter.Use(vendorMiddleware.RegisteredVendor)
	productRouter.Post("/", productController.CreateProduct)
	productRouter.Patch("/:id", productController.UpdateProduct)
	productRouter.Delete("/:id", productController.DeleteProduct)
}
