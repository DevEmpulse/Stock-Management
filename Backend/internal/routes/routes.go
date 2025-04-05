// internal/routes/routes.go
package routes

import (
	controllers "github.com/DevEmpulse/Stock-Management.git/internal/controllers/user"
	"github.com/DevEmpulse/Stock-Management.git/internal/routes/categoria"
	"github.com/DevEmpulse/Stock-Management.git/internal/routes/movimiento"
	"github.com/DevEmpulse/Stock-Management.git/internal/routes/producto"
	user "github.com/DevEmpulse/Stock-Management.git/internal/routes/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authController *controllers.AuthController) {
	api := app.Group("/api")

	// Configurar rutas de usuarios
	user.SetupUserRoutes(api, authController)

	// Configurar otras rutas
	categoria.SetupCategoriaRoutes(api)
	movimiento.SetupMovimientoRoutes(api)
	producto.SetupProductoRoutes(api)
}
