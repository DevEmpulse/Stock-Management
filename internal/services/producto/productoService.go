package services

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/productos"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/producto"
)

func GetAllProductosByIdUserService(user uint) ([]models.Productos, error) {
	return repositories.GetAllProductosByIdUser(user)
}

func CreateProducto(newProducto *models.Productos) error {
	return repositories.CreateProducto(newProducto)
}

func UpdateProducto(id uint, updatedProducto *models.Productos) error {
	return repositories.UpdateProducto(id, updatedProducto)
}

func DeleteProducto(id uint) error {
	return repositories.DeleteProducto(id)
}
