package repositories

import (
	//"fmt"

	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/almacen"
)

func GetAllAlmacenes() ([]models.Almacenes, error) {
	var almacenes []models.Almacenes

	// Buscar en la base de datos
	result := database.DB.Find(&almacenes)
	if result.Error != nil {
		return nil, result.Error
	}
	return almacenes, nil
}
