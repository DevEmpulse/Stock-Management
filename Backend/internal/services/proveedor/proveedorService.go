package services

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/proveedores"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/proveedor"
)

func CreateProveedor(newProveedor *models.Proveedores) error {
	return repositories.CreateProveedor(newProveedor)
}

func DeleteProveedor(id uint) error {
	return repositories.DeleteProveedor(id)
}

func UpdateProveedor(id uint, updatedProveedor *models.Proveedores) error {
	return repositories.UpdateProveedor(id, updatedProveedor)
}

func GetAllProveedoresByIdUser(user uint) ([]models.Proveedores, error) {
	return repositories.GetAllProveedoresByIdUser(user)
}