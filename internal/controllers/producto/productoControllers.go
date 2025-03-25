package Producto

import (
	"strconv"

	models "github.com/DevEmpulse/Stock-Management.git/internal/models/productos"
	services "github.com/DevEmpulse/Stock-Management.git/internal/services/producto"
	"github.com/gofiber/fiber/v2"
)

func GetAllProductosByIdUser(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	Productos, err := services.GetAllProductosByIdUserService(uint(idUser))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener los Productos de este usuario",
		})
	}
	return c.JSON(Productos)
}

func CreateProducto(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	var newProducto models.Productos

	if err := c.BodyParser(&newProducto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}
	newProducto.ID_users = uint(idUser)

	if err := services.CreateProducto(&newProducto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo crear la Producto",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newProducto)
}

func UpdateProducto(c *fiber.Ctx) error {
	idProductoParams := c.Params("id_producto")

	id, err := strconv.ParseInt(idProductoParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Id del Producto inválido"})
	}

	var updateProducto models.Productos

	if err := c.BodyParser(&updateProducto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}

	if err := services.UpdateProducto(uint(id), &updateProducto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar la Producto",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Producto actualizado correctamente",
	})
}

func DeleteProducto(c *fiber.Ctx) error {
	idProductoParams := c.Params("id_prod")

	id, err := strconv.ParseUint(idProductoParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID Producto inválido",
		})
	}

	if err := services.DeleteProducto(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo borrar el Producto",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Producto borrado correctamente",
	})
}
