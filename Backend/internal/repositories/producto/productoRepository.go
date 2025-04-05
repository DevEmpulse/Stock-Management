package repositories

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/productos"
)

func CreateProducto(newProducto *models.Productos) error {
	result := database.DB.Create(newProducto) // que al crear un producto cree un movimiento
	return result.Error
}

func GetAllProductosByIdUser(user uint) ([]models.Productos, error) {
	var ProductosUser []models.Productos

	result := database.DB.
		Preload("User").
		Preload("Categoria"). 
		Where("id_users = ?", user).
		Find(&ProductosUser)

	if result.Error != nil {
		return nil, result.Error
	}
	return ProductosUser, nil
}

func DeleteProducto(id uint) error {
	result := database.DB.Delete(&models.Productos{}, id)
	return result.Error
}

// Actualizar un Producto por ID
func UpdateProducto(id uint, updatedProducto *models.Productos) error {
	var existeProducto models.Productos
	result := database.DB.First(&existeProducto, id)
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Model(&existeProducto).Updates(updatedProducto) // que al agregar mas productos cree un movimiento
	return result.Error
}
