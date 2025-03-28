package almacen

import (
	"strconv"

	models "github.com/DevEmpulse/Stock-Management.git/internal/models/almacen"
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

func GetAllAlmacenesByIdUser(c *fiber.Ctx) error {

	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de usuario invalido"})
	}
	almacenes, err := services.GetAllAlmacenesByIdUserService(uint(idUser))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener los almacenes de este usuario",
		})
	}
	return c.JSON(almacenes)
}

func CreateAlmacen(c *fiber.Ctx) error {

	idUserParam := c.Params("id_user")

	// Convertir el ID de string a uint
	idUser, err := strconv.ParseUint(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID de usuario inválido"})
	}

	var newAlmacen models.Almacenes

	if err := c.BodyParser(&newAlmacen); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}

	newAlmacen.ID_users = uint(idUser)

	if err := services.CreateAlmacen(&newAlmacen); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo crear el Almacen",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newAlmacen)
}

func UpdateAlmacen(c *fiber.Ctx) error {

	idAlmacenParams := c.Params("id_almacen")

	id, err := strconv.ParseInt(idAlmacenParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID del almacen inválido"})
	}

	var updateAlmacen models.Almacenes

	if err := c.BodyParser(&updateAlmacen); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}

	if err := services.UpdateAlmacenService(uint(id), &updateAlmacen); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar el almacén",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Almacén actualizado correctamente",
	})
}

func DeleteAlmacen(c *fiber.Ctx) error {
	idAlmacenParams := c.Params("id_almacen")

	id, err := strconv.ParseUint(idAlmacenParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID del almacen inválido"})
	}

	if err := services.DeleteAlmacen(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Nos se pudo borrar este almacen"})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Almacén Borrado correctamente",
	})

}
