package repositories

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/almacen"
)

func GetAllAlmacenes() ([]models.Almacenes, error) {
	var almacenes []models.Almacenes

	// Precargar la relación User
	result := database.DB.Preload("User").Find(&almacenes)
	if result.Error != nil {
		return nil, result.Error
	}
	return almacenes, nil
}

func GetAllAlmacenesByIdUser(user uint) ([]models.Almacenes, error) {
	var almacenesUsers []models.Almacenes

	// Precargar la relación User
	result := database.DB.Preload("User").Where("id_users = ?", user).Find(&almacenesUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	return almacenesUsers, nil
}

func CreateAlmacen(newAlmacen *models.Almacenes) error {
	result := database.DB.Create(newAlmacen)

	return result.Error

}

func UpdateAlmacen(id uint, updatedAlmacen *models.Almacenes) error {
	var existingAlmacen models.Almacenes
	result := database.DB.First(&existingAlmacen, id)
	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Model(&existingAlmacen).Updates(updatedAlmacen)
	return result.Error
}

func DeleteAlmacen(id uint) error {

	result := database.DB.Delete(&models.Almacenes{}, id)
	return result.Error
}
