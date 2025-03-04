package models

type Users struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre"`
	Email  string `json:"email" gorm:"unique"`
}
