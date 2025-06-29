package models

import "time"

type Expediente struct {
	ID                     uint      `json:"id" gorm:"primaryKey"`
	PacienteID             uint      `json:"paciente_id" gorm:"unique"`
	GrupoSanguineo         string    `json:"grupo_sanguineo"`
	Alergias               string    `json:"alergias"`
	EnfermedadesCronicas   string    `json:"enfermedades_cronicas"`
	AntecedentesFamiliares string    `json:"antecedentes_familiares"`
	AntecedentesPersonales string    `json:"antecedentes_personales"`
	MedicamentosHabituales string    `json:"medicamentos_habituales"`
	Vacunas                string    `json:"vacunas"`
	NotasGenerales         string    `json:"notas_generales"`
	FechaActualizacion     time.Time `json:"fecha_actualizacion" gorm:"type:date"`
}
