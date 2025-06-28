package models

import (
	"time"
)

// models/usuario.go
type Usuario struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Nombre          string    `json:"nombre"`
	Apellidos       string    `json:"apellidos"`
	Tipo            string    `json:"tipo"`
	FechaNacimiento time.Time `json:"fecha_nacimiento" gorm:"type:date" time_format:"2006-01-02"`
	Genero          string    `json:"genero"`
	Correo          string    `json:"correo"`
	Password        string    `json:"-"` // no exponer
	CreadoEn        time.Time `json:"creado_en"`
	ActualizadoEn   time.Time `json:"actualizado_en"`
}
