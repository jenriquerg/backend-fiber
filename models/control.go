package models

import "time"

type Control struct {
	ID                   uint      `json:"id" gorm:"primaryKey"`
	PacienteID           uint      `json:"paciente_id"`
	PesoKg               float64   `json:"peso_kg"`
	AlturaCm             float64   `json:"altura_cm"`
	IMC                  float64   `json:"imc"`
	PresionArterial      string    `json:"presion_arterial"`
	FrecuenciaCardiaca   int       `json:"frecuencia_cardiaca"`
	FrecuenciaRespiratoria int     `json:"frecuencia_respiratoria"`
	TemperaturaC         float64   `json:"temperatura_c"`
	NivelGlucosa         float64   `json:"nivel_glucosa"`
	SaturacionOxigeno    float64   `json:"saturacion_oxigeno"`
	NotasGenerales       string    `json:"notas_generales"`
	Fecha                time.Time `json:"fecha" gorm:"type:date"`
}

func (Control) TableName() string {
	return "controles"
}
