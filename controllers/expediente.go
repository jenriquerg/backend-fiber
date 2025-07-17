package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
	"time"
)

func GetExpedientes(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "E01")
	c.Locals("intCodeError", "F17")
	var expedientes []models.Expediente
	if err := config.DB.Find(&expedientes).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener expedientes"})
	}
	return c.JSON(expedientes)
}

func GetExpediente(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "E02")
	c.Locals("intCodeError", "F18")
	id := c.Params("paciente_id")
	var expediente models.Expediente
	if err := config.DB.Where("paciente_id = ?", id).First(&expediente).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Expediente no encontrado"})
	}
	return c.JSON(expediente)
}

func CreateExpediente(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "E03")
	c.Locals("intCodeError", "F19")
	type rawExpediente struct {
		PacienteID             uint   `json:"paciente_id"`
		GrupoSanguineo         string `json:"grupo_sanguineo"`
		Alergias               string `json:"alergias"`
		EnfermedadesCronicas   string `json:"enfermedades_cronicas"`
		AntecedentesFamiliares string `json:"antecedentes_familiares"`
		AntecedentesPersonales string `json:"antecedentes_personales"`
		MedicamentosHabituales string `json:"medicamentos_habituales"`
		Vacunas                string `json:"vacunas"`
		NotasGenerales         string `json:"notas_generales"`
		FechaActualizacion     string `json:"fecha_actualizacion"` // YYYY-MM-DD
	}

	var input rawExpediente
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	fecha, err := time.Parse("2006-01-02", input.FechaActualizacion)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Fecha inválida. Usa YYYY-MM-DD"})
	}

	exp := models.Expediente{
		PacienteID:             input.PacienteID,
		GrupoSanguineo:         input.GrupoSanguineo,
		Alergias:               input.Alergias,
		EnfermedadesCronicas:   input.EnfermedadesCronicas,
		AntecedentesFamiliares: input.AntecedentesFamiliares,
		AntecedentesPersonales: input.AntecedentesPersonales,
		MedicamentosHabituales: input.MedicamentosHabituales,
		Vacunas:                input.Vacunas,
		NotasGenerales:         input.NotasGenerales,
		FechaActualizacion:     fecha,
	}

	if err := config.DB.Create(&exp).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo crear expediente"})
	}
	return c.Status(201).JSON(exp)
}

func UpdateExpediente(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "E04")
	c.Locals("intCodeError", "F20")
	id := c.Params("id")

	var input map[string]interface{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if dateStr, ok := input["fecha_actualizacion"].(string); ok && dateStr != "" {
		fecha, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Fecha inválida"})
		}
		input["fecha_actualizacion"] = fecha
	}

	var expediente models.Expediente
	if err := config.DB.First(&expediente, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Expediente no encontrado"})
	}

	if err := config.DB.Model(&expediente).Updates(input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo actualizar expediente"})
	}
	return c.JSON(expediente)
}

func DeleteExpediente(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "E05")
	c.Locals("intCodeError", "F21")
	id := c.Params("id")
	if err := config.DB.Delete(&models.Expediente{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo eliminar expediente"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Expediente eliminado"})
}
