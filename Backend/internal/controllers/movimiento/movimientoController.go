package Movimiento

import (
	"strconv"

	models "github.com/DevEmpulse/Stock-Management.git/internal/models/movimientos"
	services "github.com/DevEmpulse/Stock-Management.git/internal/services/movimiento"
	"github.com/gofiber/fiber/v2"
)

func GetAllMovimientosByIdUser(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	movimientos, err := services.GetAllMovimientosByIdUserService(uint(idUser))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener los Movimientos de este usuario",
		})
	}
	return c.JSON(movimientos)
}

func GetEgresosByIdUser(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario inválido",
		})
	}

	egresos, err := services.GetEgresosByIdUser(uint(idUser))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener los egresos de este usuario",
		})
	}

	return c.JSON(egresos)
}

func GetIngresosByIdUser(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario inválido",
		})
	}

	ingresos, err := services.GetIngresosByIdUser(uint(idUser))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener los ingresos de este usuario",
		})
	}

	return c.JSON(ingresos)
}

func CreateMovimiento(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	var newMovimiento models.Movimientos

	if err := c.BodyParser(&newMovimiento); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}
	newMovimiento.ID_users = uint(idUser)

	if err := services.CreateMovimiento(&newMovimiento); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo crear el Movimiento",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newMovimiento)
}

func UpdateMovimiento(c *fiber.Ctx) error {
	idMovimientoParams := c.Params("id_movimiento")

	id, err := strconv.ParseInt(idMovimientoParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Id del Movimiento inválido"})
	}

	var updateMovimiento models.Movimientos

	if err := c.BodyParser(&updateMovimiento); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}

	if err := services.UpdateMovimiento(uint(id), &updateMovimiento); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar la Movimiento",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Movimiento actualizado correctamente",
	})
}

func DeleteMovimiento(c *fiber.Ctx) error {
	idMovimientoParams := c.Params("id_movimiento")

	id, err := strconv.ParseUint(idMovimientoParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID movimiento inválido",
		})
	}

	if err := services.DeleteMovimiento(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo borrar la Movimiento",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Movimiento borrado correctamente",
	})
}
