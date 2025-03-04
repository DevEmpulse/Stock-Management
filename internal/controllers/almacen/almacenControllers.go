package almacen

import (
	services "github.com/DevEmpulse/Stock-Management.git/internal/services/almacen"
	"github.com/gofiber/fiber/v2"
)

// Controlador para manejar la petición GET de almacenes
func GetAllAlmacenes(c *fiber.Ctx) error {
	almacenes, err := services.GetAllAlmacenesService()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener los almacenes",
		})
	}

	return c.JSON(almacenes)
}
