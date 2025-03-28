package models

type Users struct {
	ID_users   uint   `json:"id_users" gorm:"primaryKey"`
	Nombre     string `json:"nombre"`
	Nickname   string `json:"nickname" gorm:"unique"`
	Logo       string `json:"logo"`
	Email      string `json:"email" gorm:"unique"`
	Contrasena string `json:"-"` // Excluir de la respuesta JSON
}
