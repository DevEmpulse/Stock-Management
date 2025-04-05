package models

import (
	productos "github.com/DevEmpulse/Stock-Management.git/internal/models/productos"
	models "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
)

type Movimientos struct {
	ID_mov      uint                `json:"id_mov" gorm:"primaryKey"`
	ID_prod     uint                `json:"id_producto"`
	Producto    productos.Productos `json:"producto" gorm:"foreignKey:ID_prod;references:ID_prod"`
	TipoDeMov   string              `json:"tipo_de_mov"`
	Monto 		float64             `json:"monto"`
	FechaMov    string              `json:"fecha_movimiento"`
	ID_users    uint                `json:"id_users"`
	User        models.Users        `json:"user" gorm:"foreignKey:ID_users;references:ID_users"`
	Cantidad    int                 `json:"cantidad"`
}
