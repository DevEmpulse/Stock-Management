package services

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/almacen"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/almacen"
)

// Servicio para obtener todos los almacenes
func GetAllAlmacenesService() ([]models.Almacen, error) {
	return repositories.GetAllAlmacenes()
}
