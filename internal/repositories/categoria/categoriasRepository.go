package repositories

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/categorias"
)

func CreateCategoria(newCategoria *models.Categorias) error {
	result := database.DB.Create(newCategoria)

	return result.Error

}

func GetAllCategoriasByIdUser(user uint) ([]models.Categorias, error) {
	var categoriasUser []models.Categorias

	result := database.DB.Preload("User").Where("id_users = ?", user).Find(&categoriasUser)

	if result.Error != nil {
		return nil, result.Error
	}
	return categoriasUser, nil
}

func DeleteCategoria(id uint) error {

	result := database.DB.Delete(&models.Categorias{}, id)
	return result.Error

}

func UpdateCategoria(id uint, updatedCategoria *models.Categorias) error {
	var existeCategoria models.Categorias
	result := database.DB.First(&existeCategoria, id)
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Model(&existeCategoria).Updates(updatedCategoria)
	return result.Error
}
