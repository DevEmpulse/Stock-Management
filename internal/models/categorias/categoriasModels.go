package models

import (
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
)

type Categorias struct {
	ID_categoria uint         `json:"id_categoria" gorm:"primaryKey"`
	ID_users     uint         `json:"id_users"`
	User         models.Users `json:"user" gorm:"foreignKey:ID_users;references:ID_users"`
	Nombre       string       `json:"nombre"`
}
