package models

import "time"

type Consulta struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	IDConsultorio uint      `json:"id_consultorio"`
	IDMedico     uint      `json:"id_medico"`
	IDPaciente   uint      `json:"id_paciente"`
	Tipo         string    `json:"tipo"`
	Horario      time.Time `json:"horario"`
	Diagnostico  string    `json:"diagnostico"`
}