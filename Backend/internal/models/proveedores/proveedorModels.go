package models

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
)

type Proveedores struct {
	ID_proveedor     uint         `json:"id_proveedor" gorm:"primaryKey"`
	CuitCuil         string       `json:"cuit_cuil"`
	Nombre_apellido  string       `json:"nombre_apellido"`
	Email            string       `json:"email"`
	NumContacto      string       `json:"num_contacto"`
	Comentario       string       `json:"comentario"`
	ID_user          uint         `json:"id_user"`
	User             models.Users `json:"user" gorm:"foreignKey:ID_user;references:ID_users"`
}