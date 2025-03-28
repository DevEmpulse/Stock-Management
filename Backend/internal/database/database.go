package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Cargar variables de entorno
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error al cargar el archivo .env:", err)
	}

	// Obtener variables de entorno
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Configuración de conexión a PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	// ✅ ASIGNAR `db` A `DB` PARA QUE SE PUEDA USAR DESDE OTROS ARCHIVOS
	DB = db

	// Verificar la conexión
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error al obtener la conexión de base de datos:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	} else {
		fmt.Println("¡Conexión exitosa a PostgreSQL!")
	}
}
