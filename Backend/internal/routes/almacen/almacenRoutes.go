package almacen

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/controllers/almacen"
	"github.com/gofiber/fiber/v2"
)

func SetupAlmacenRoutes(api fiber.Router) {
	almacenRoutes := api.Group("/almacen")

	almacenRoutes.Get("/", almacen.GetAllAlmacenes) //medio que no sirve, porqur la idea es que cada usuario vea su almacen
	almacenRoutes.Get("/:id_user", almacen.GetAllAlmacenesByIdUser)
	almacenRoutes.Post("/:id_user", almacen.CreateAlmacen)
	almacenRoutes.Put("/:id_almacen", almacen.UpdateAlmacen)
	almacenRoutes.Delete("/:id_almacen", almacen.DeleteAlmacen)
}
