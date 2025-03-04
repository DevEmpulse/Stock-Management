package routes

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/routes/almacen"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	almacen.SetupAlmacenRoutes(api)
}
