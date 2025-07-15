package models

import "time"

type Receta struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	IdConsulta  uint      `json:"id_consulta"`
	Fecha       time.Time `json:"fecha" gorm:"type:date"`
	IdMedico    uint      `json:"id_medico"`
	Medicamento string    `json:"medicamento"`
	Dosis       string    `json:"dosis"`
}

func (Receta) TableName() string {
	return "recetas"
}