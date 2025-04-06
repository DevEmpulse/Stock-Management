package proveedor

import (
	proveedor "github.com/DevEmpulse/Stock-Management.git/internal/controllers/Proveedor"
	"github.com/gofiber/fiber/v2"
)

func SetupProveedorRoutes(api fiber.Router) {
	proveedorRoutes := api.Group("/proveedor")

	proveedorRoutes.Get("/:id_user", proveedor.GetAllProveedoresByIdUser)
	proveedorRoutes.Post("/:id_user", proveedor.CreateProveedor)
	proveedorRoutes.Delete("/:id_proveedor", proveedor.DeleteProveedor)
	proveedorRoutes.Put("/:id_proveedor", proveedor.UpdateProveedor)
}