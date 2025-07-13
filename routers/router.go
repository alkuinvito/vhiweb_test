package routers

import (
	"vhiweb_test/app/users"
	"vhiweb_test/lib/adapters"
	"vhiweb_test/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Handle(app *fiber.App) {
	db := adapters.NewDB()
	userController := users.NewUserController(db)
	authMiddleware := middlewares.NewAuthMiddleware(&users.UserService{})

	app.Use(recover.New())
	app.Use(cors.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	userRouter := v1.Group("/users")
	userRouter.Get("/", userController.GetUsers)
	userRouter.Get("/:id", userController.GetUserById)
	// authenticated user only
	userRouter.Use(authMiddleware.Authenticated)
	userRouter.Patch("/:id", userController.UpdateUser)
	userRouter.Delete("/:id", userController.DeleteUser)

	authRouter := v1.Group("/auth")
	authRouter.Post("/login", userController.Login)
	authRouter.Post("/register", userController.Register)
}
