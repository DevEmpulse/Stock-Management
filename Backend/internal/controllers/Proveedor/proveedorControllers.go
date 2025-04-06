package proveedor

import (
	"strconv"

	models "github.com/DevEmpulse/Stock-Management.git/internal/models/proveedores"
	services "github.com/DevEmpulse/Stock-Management.git/internal/services/proveedor"
	"github.com/gofiber/fiber/v2"
)

func GetAllProveedoresByIdUser(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	Proveedores, err := services.GetAllProveedoresByIdUser(uint(idUser))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener las proveedores de este usuario",
		})
	}
	return c.JSON(Proveedores)
}

func CreateProveedor(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	var newProveedor models.Proveedores

	if err := c.BodyParser(&newProveedor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}
	newProveedor.ID_user = uint(idUser)

	if err := services.CreateProveedor(&newProveedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo crear el proveedor",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newProveedor)
}

func DeleteProveedor(c *fiber.Ctx) error {
	idProveedorParams := c.Params("id_proveedor")

	id, err := strconv.ParseUint(idProveedorParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID categoria inválido",
		})
	}

	if err := services.DeleteProveedor(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo borrar el proveedor",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proveedor borrada correctamente",
	})
}

func UpdateProveedor(c *fiber.Ctx) error {
	idProveedorParams := c.Params("id_proveedor")

	id, err := strconv.ParseInt(idProveedorParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Id del proveedor inválido"})
	}

	var UpdateProveedor models.Proveedores

	if err := c.BodyParser(&UpdateProveedor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}

	if err := services.UpdateProveedor(uint(id), &UpdateProveedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar el Proveedor",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proveedor actualizado correctamente",
	})
}