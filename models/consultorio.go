package models

type Consultorio struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Nombre    string `json:"nombre"`
	Tipo      string `json:"tipo"`
	IDMedico  *uint  `json:"id_medico"`
	Status    string `json:"status"`
	Ubicacion string `json:"ubicacion"`
	Horario   string `json:"horario"`
}