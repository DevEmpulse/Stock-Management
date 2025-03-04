package almacen

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/controllers/almacen"
	"github.com/gofiber/fiber/v2"
)

func SetupAlmacenRoutes(api fiber.Router) {
	almacenRoutes := api.Group("/almacen")

	almacenRoutes.Get("/", almacen.GetAllAlmacenes)

}
