package routers

import (
	"vhiweb_test/app/users"
	"vhiweb_test/lib/adapters"

	"github.com/gofiber/fiber/v2"
)

func Handle(app *fiber.App) {
	db := adapters.NewDB()
	userController := users.NewUserController(db)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	userRouter := v1.Group("/users")
	userRouter.Get("/", userController.GetUsers)
	userRouter.Get("/:id", userController.FindUserById)
	userRouter.Put("/:id", userController.UpdateUser)
	userRouter.Delete("/:id", userController.DeleteUser)

	authRouter := v1.Group("/auth")
	authRouter.Post("/login", userController.Login)
	authRouter.Post("/register", userController.Register)
}
