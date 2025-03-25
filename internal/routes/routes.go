package routes

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/routes/almacen"
	"github.com/DevEmpulse/Stock-Management.git/internal/routes/categoria"
	"github.com/DevEmpulse/Stock-Management.git/internal/routes/movimiento"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	almacen.SetupAlmacenRoutes(api)
	categoria.SetupCategoriaRoutes(api)
	movimiento.SetupMovimientoRoutes(api)
}
