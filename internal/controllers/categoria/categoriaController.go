package categoria

import (
	"strconv"

	models "github.com/DevEmpulse/Stock-Management.git/internal/models/categorias"
	services "github.com/DevEmpulse/Stock-Management.git/internal/services/categoria"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategoriasByIdUser(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	categorias, err := services.GetAllCategoriasByIdUserService(uint(idUser))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error al obtener las categorias de este usuario",
		})
	}
	return c.JSON(categorias)
}

func CreateCategoria(c *fiber.Ctx) error {
	idUserParam := c.Params("id_user")

	idUser, err := strconv.ParseInt(idUserParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de Usuario invalido"})
	}

	var newCategoria models.Categorias

	if err := c.BodyParser(&newCategoria); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}
	newCategoria.ID_users = uint(idUser)

	if err := services.CreateCategoria(&newCategoria); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo crear la categoria",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newCategoria)
}

func UpdateCategoria(c *fiber.Ctx) error {
	idCategoriaParams := c.Params("id_categoria")

	id, err := strconv.ParseInt(idCategoriaParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Id de la categoria inválido"})
	}

	var updateCategoria models.Categorias

	if err := c.BodyParser(&updateCategoria); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos ingresados invalidos",
		})
	}

	if err := services.UpdateCategoria(uint(id), &updateCategoria); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar la categoria",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Categoria actualizado correctamente",
	})
}

func DeleteCategoria(c *fiber.Ctx) error {
	idCategoriaParams := c.Params("id_categoria")

	id, err := strconv.ParseUint(idCategoriaParams, 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID almacen inválido",
		})
	}

	if err := services.DeleteCategoria(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo borrar la categoria",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Categoria borrada correctamente",
	})
}
