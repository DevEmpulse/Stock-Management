package producto

import (
	producto "github.com/DevEmpulse/Stock-Management.git/internal/controllers/producto"
	"github.com/gofiber/fiber/v2"
)

func SetupProductoRoutes(api fiber.Router) {
	productoRoutes := api.Group("/producto")

	productoRoutes.Get("/:id_user", producto.GetAllProductosByIdUser)
	productoRoutes.Post("/:id_user", producto.CreateProducto)
	productoRoutes.Put("/:id_producto", producto.UpdateProducto)
	productoRoutes.Delete("/:id_producto", producto.DeleteProducto)

}
