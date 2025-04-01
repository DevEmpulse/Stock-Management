package controllers

import (
	"strconv"

	users "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
	services "github.com/DevEmpulse/Stock-Management.git/internal/services/user"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	user := new(users.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := ac.AuthService.Register(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.JSON(user)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := ac.AuthService.Login(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (ac *AuthController) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := ac.AuthService.DeleteUser(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error deleting user"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User deleted successfully"})
}

func (ac *AuthController) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var updatedUser users.Users
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := ac.AuthService.UpdateUser(uint(id), &updatedUser); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error updating user"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User updated successfully"})
}
