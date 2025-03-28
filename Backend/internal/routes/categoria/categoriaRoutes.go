package categoria

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/controllers/categoria"
	"github.com/gofiber/fiber/v2"
)

func SetupCategoriaRoutes(api fiber.Router) {
	categoriaRoutes := api.Group("/categoria")

	categoriaRoutes.Get("/:id_user", categoria.GetAllCategoriasByIdUser)
	categoriaRoutes.Post("/:id_user", categoria.CreateCategoria)
	categoriaRoutes.Put("/:id_categoria", categoria.UpdateCategoria)
	categoriaRoutes.Delete("/:id_categoria", categoria.DeleteCategoria)

}
