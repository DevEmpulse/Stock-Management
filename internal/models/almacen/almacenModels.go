package models

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
)

type Almacenes struct {
	ID_almacen uint         `json:"id_almacen" gorm:"primaryKey"`
	ID_users   uint         `json:"id_users"`
	User       models.Users `json:"user" gorm:"foreignKey:ID_users;references:ID"`
	Nombre     string       `json:"nombre"`
	Direccion  string       `json:"direccion"`
}
