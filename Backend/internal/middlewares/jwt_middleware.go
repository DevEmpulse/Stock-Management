package middlewares

import (
	"os"

	"github.com/dgrijalva/jwt-go" // Importa jwt-go
	"github.com/gofiber/fiber/v2"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["id"].(float64)) // ⚠️ importante: convertir de float64 a uint

		c.Locals("user", userID)
		return c.Next()
	}
}
