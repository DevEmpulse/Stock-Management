package repositories

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/almacen"
)

func GetAllAlmacenes() ([]models.Almacen, error) {
	var almacenes []models.Almacen

	// Buscar en la base de datos
	result := database.DB.Find(&almacenes)
	if result.Error != nil {
		return nil, result.Error
	}

	return almacenes, nil
}
