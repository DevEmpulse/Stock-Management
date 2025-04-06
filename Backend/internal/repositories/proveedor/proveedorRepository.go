package repositories

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/proveedores"
)

func CreateProveedor (newProveedor *models.Proveedores) error {
	result:= database.DB.Create(newProveedor)
	return result.Error
}

func DeleteProveedor(id uint) error {
	result := database.DB.Delete(&models.Proveedores{}, id)
	return result.Error
}

func UpdateProveedor(id uint, updatedProveedor *models.Proveedores) error {
	var existeProveedor models.Proveedores
	result := database.DB.First(&existeProveedor, id)
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Model(&existeProveedor).Updates(updatedProveedor)
	return result.Error
}

func GetAllProveedoresByIdUser(user uint) ([]models.Proveedores, error) {
	var proveedoresUser []models.Proveedores

	result := database.DB.Preload("User").Where("id_user = ?", user).Find(&proveedoresUser)

	if result.Error != nil {
		return nil, result.Error
	}
	return proveedoresUser, nil
}