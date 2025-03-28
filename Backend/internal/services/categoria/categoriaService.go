package services

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/categorias"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/categoria"
)

func GetAllCategoriasByIdUserService(user uint) ([]models.Categorias, error) {
	return repositories.GetAllCategoriasByIdUser(user)
}

func CreateCategoria(newCategoria *models.Categorias) error {
	return repositories.CreateCategoria(newCategoria)
}

func UpdateCategoria(id uint, updatedCategoria *models.Categorias) error {
	return repositories.UpdateCategoria(id, updatedCategoria)
}

func DeleteCategoria(id uint) error {
	return repositories.DeleteCategoria(id)
}
