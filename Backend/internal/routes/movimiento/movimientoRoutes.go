package movimiento

import (
	Movimiento "github.com/DevEmpulse/Stock-Management.git/internal/controllers/movimiento"
	"github.com/gofiber/fiber/v2"
)

func SetupMovimientoRoutes(api fiber.Router) {
	MovimientoRoutes := api.Group("/movimiento")

	MovimientoRoutes.Get("/ingresos/:id_user", Movimiento.GetIngresosByIdUser)
	MovimientoRoutes.Get("/egresos/:id_user", Movimiento.GetEgresosByIdUser)
	MovimientoRoutes.Get("/:id_user", Movimiento.GetAllMovimientosByIdUser)
	MovimientoRoutes.Post("/:id_user", Movimiento.CreateMovimiento)
	MovimientoRoutes.Put("/:id_movimiento", Movimiento.UpdateMovimiento)
	MovimientoRoutes.Delete("/:id_movimiento", Movimiento.DeleteMovimiento)

}
