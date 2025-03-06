package services

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/almacen"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/almacen"
)

// Servicio para obtener todos los almacenes
func GetAllAlmacenesService() ([]models.Almacenes, error) {
	return repositories.GetAllAlmacenes()
}

func GetAllAlmacenesByIdUserService(user uint) ([]models.Almacenes, error) {
	return repositories.GetAllAlmacenesByIdUser(user)
}

func CreateAlmacen(newAlmacen *models.Almacenes) error {
	return repositories.CreateAlmacen(newAlmacen)
}

func UpdateAlmacenService(id uint, updatedAlmacen *models.Almacenes) error {
	return repositories.UpdateAlmacen(id, updatedAlmacen)
}

func DeleteAlmacen(id uint) error {
	return repositories.DeleteAlmacen(id)
}
