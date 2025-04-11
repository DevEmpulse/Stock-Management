package routes

import (
	controllers "github.com/DevEmpulse/Stock-Management.git/internal/controllers/user"
	"github.com/DevEmpulse/Stock-Management.git/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, authController *controllers.AuthController) {
	userRoutes := api.Group("/users")

	userRoutes.Post("/register", authController.Register)
	userRoutes.Post("/login", authController.Login)
	userRoutes.Get("/getUser", middlewares.JWTMiddleware(), authController.GetUser)
	userRoutes.Put("/:id", middlewares.JWTMiddleware(), authController.UpdateUser)
	userRoutes.Delete("/:id", middlewares.JWTMiddleware(), authController.DeleteUser)
}
