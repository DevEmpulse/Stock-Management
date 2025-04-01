package main

import (
	"log"

	controllers "github.com/DevEmpulse/Stock-Management.git/internal/controllers/user"
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/user"
	"github.com/DevEmpulse/Stock-Management.git/internal/routes"
	services "github.com/DevEmpulse/Stock-Management.git/internal/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Conectar a la base de datos
	database.Connect()

	// Configurar dependencias
	userRepo := repositories.NewUserRepository()                 // Crear repositorio
	authService := services.NewAuthService(userRepo)             // Crear servicio con el repositorio
	authController := controllers.NewAuthController(authService) // Crear controlador con el servicio

	// Configuración de Fiber
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Cambia esto al dominio de tu frontend
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Authorization",
	}))
	// Configurar rutas con el controlador
	routes.SetupRoutes(app, authController)

	// Iniciar el servidor
	log.Fatal(app.Listen(":3000"))
}
