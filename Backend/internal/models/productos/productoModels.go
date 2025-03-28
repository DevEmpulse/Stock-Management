package models

import (
	categoria "github.com/DevEmpulse/Stock-Management.git/internal/models/categorias"
	users "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
)

// FALTA LO DE PROVEEDOR ACA RELACIONADO A LO DE MODELS
type Productos struct {
	ID_prod        uint                 `json:"id_prod" gorm:"primaryKey"`
	CodigoDeBarra  string               `json:"codigo_de_barra"`
	ID_almacen     uint                 `json:"id_almacen"`
	ID_users       uint                 `json:"id_users"`
	User           users.Users          `json:"user" gorm:"foreignKey:ID_users;references:ID_users"`
	Nombre         string               `json:"nombre"`
	Imagen         string               `json:"imagen"`
	Stock          int                  `json:"stock"`
	ID_categoria   uint                 `json:"id_categoria"`
	Categoria      categoria.Categorias `json:"categoria" gorm:"foreignKey:ID_categoria;references:ID_categoria"`
	ID_proveedor   uint                 `json:"id_proveedor"`
	Descripcion    string               `json:"descripcion"`
	PrecioCompra   float64              `json:"precio_compra"`
	PrecioVenta    float64              `json:"precio_venta"`
	FechaCreacion  string               `json:"fecha_creacion"`
	FechaActualiza string               `json:"fecha_actualizacion"`
}
