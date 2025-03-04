package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
    DB *gorm.DB
}

// Rutas
func (r *Repository) SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
    _ = api // Solo para evitar que el compilador marque esto como no usado
}

func main() {
    // Cargar variables de entorno
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error al cargar el archivo .env:", err)
    }

    // Obtener las variables de entorno
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    // Configuración de la base de datos con las variables de entorno
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("No se pudo conectar a la base de datos:", err)
    }

    // Obtener la conexión *sql.DB de GORM para hacer ping
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("Error al obtener la conexión de base de datos:", err)
    }

    // Verificar la conexión con Ping
    err = sqlDB.Ping()
    if err != nil {
        log.Fatal("No se pudo conectar a la base de datos:", err)
    } else {
        fmt.Println("¡Conexión exitosa a PostgreSQL!")
    }

    // Configuración de Fiber
    r := Repository{DB: db}
    app := fiber.New()
    r.SetupRoutes(app)
    log.Fatal(app.Listen(":3000"))
}
