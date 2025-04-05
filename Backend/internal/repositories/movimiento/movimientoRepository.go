package repositories

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/movimientos"
)

func CreateMovimiento(newMovimiento *models.Movimientos) error {
	result := database.DB.Create(newMovimiento)
	return result.Error
}

func GetAllMovimientosByIdUser(user uint) ([]models.Movimientos, error) {
	var movimientosUser []models.Movimientos

	result := database.DB.
		Preload("User").
		Preload("Producto").
		Where("id_users = ?", user).
		Find(&movimientosUser)

	if result.Error != nil {
		return nil, result.Error
	}
	return movimientosUser, nil
}

func GetEgresosByIdUser(userID uint) ([]models.Movimientos, error) {
	var egresos []models.Movimientos

	result := database.DB.
		Preload("User").
		Preload("Producto").
		Where("id_users = ? AND tipo_de_mov = ?", userID, "egreso").
		Find(&egresos)

	if result.Error != nil {
		return nil, result.Error
	}
	return egresos, nil
}

func GetIngresosByIdUser(userID uint) ([]models.Movimientos, error) {
	var ingresos []models.Movimientos

	result := database.DB.
		Preload("User").
		Preload("Producto").
		Where("id_users = ? AND tipo_de_mov = ?", userID, "ingreso").
		Find(&ingresos)

	if result.Error != nil {
		return nil, result.Error
	}
	return ingresos, nil
}

func DeleteMovimiento(id uint) error {
	result := database.DB.Delete(&models.Movimientos{}, id)
	return result.Error
}

// Actualizar un movimiento por ID
func UpdateMovimiento(id uint, updatedMovimiento *models.Movimientos) error {
	var existeMovimiento models.Movimientos
	result := database.DB.First(&existeMovimiento, id)
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Model(&existeMovimiento).Updates(updatedMovimiento)
	return result.Error
}
