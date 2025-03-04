package main

import (
	"log"

	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	"github.com/DevEmpulse/Stock-Management.git/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	// Configuración de Fiber
	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
