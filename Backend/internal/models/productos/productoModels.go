package models

import (
	"time"
)

// FALTA LO DE PROVEEDOR ACA RELACIONADO A LO DE MODELS
type Productos struct {
	ID_prod        uint      `json:"id_prod" gorm:"primaryKey"`
	CodigoDeBarra  string    `json:"codigo_de_barras" gorm:"column:codigo_de_barras"`
	ID_users       uint      `json:"id_user" gorm:"column:id_user"`
	Nombre         string    `json:"nombre"`
	Imagen         string    `json:"imagen"`
	Stock          int       `json:"stock"`
	ID_categoria   uint      `json:"id_categoria"`
	Descripcion    string    `json:"descripcion"`
	PrecioCompra   float64   `json:"precio_compra"`
	PrecioVenta    float64   `json:"precio_venta"`
	FechaCreacion  time.Time `json:"fecha_creacion" gorm:"column:fecha_de_creacion"`
	FechaActualiza time.Time `json:"fecha_actualizacion" gorm:"column:fecha_de_actualizacion"`
}
