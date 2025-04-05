package services

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/movimientos"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/movimiento"
)

func GetAllMovimientosByIdUserService(user uint) ([]models.Movimientos, error) {
	return repositories.GetAllMovimientosByIdUser(user)
}

func GetCompraByIdUser(user uint) ([]models.Movimientos, error) {
	return repositories.GetAllMovimientosByIdUser(user)
}

func GetVentaByIdUser(user uint) ([]models.Movimientos, error) {
	return repositories.GetAllMovimientosByIdUser(user)
}

func CreateMovimiento(newMovimiento *models.Movimientos) error {
	return repositories.CreateMovimiento(newMovimiento)
}

func UpdateMovimiento(id uint, updatedMovimiento *models.Movimientos) error {
	return repositories.UpdateMovimiento(id, updatedMovimiento)
}

func DeleteMovimiento(id uint) error {
	return repositories.DeleteMovimiento(id)
}
