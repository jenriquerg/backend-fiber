package controllers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
)

func GetControles(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CN01")
	c.Locals("intCodeError", "F13")
	var controles []models.Control
	if err := config.DB.Find(&controles).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener controles"})
	}
	return c.JSON(controles)
}

func GetControl(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CN02")
	c.Locals("intCodeError", "F14")
	id := c.Params("id")
	var control models.Control
	if err := config.DB.First(&control, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Control no encontrado"})
	}
	return c.JSON(control)
}

func CreateControl(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CN03")
	c.Locals("intCodeError", "F15")
	type rawControl struct {
		PacienteID             uint    `json:"paciente_id"`
		PesoKg                 float64 `json:"peso_kg"`
		AlturaCm               float64 `json:"altura_cm"`
		IMC                    float64 `json:"imc"`
		PresionArterial        string  `json:"presion_arterial"`
		FrecuenciaCardiaca     int     `json:"frecuencia_cardiaca"`
		FrecuenciaRespiratoria int     `json:"frecuencia_respiratoria"`
		TemperaturaC           float64 `json:"temperatura_c"`
		NivelGlucosa           float64 `json:"nivel_glucosa"`
		SaturacionOxigeno      float64 `json:"saturacion_oxigeno"`
		NotasGenerales         string  `json:"notas_generales"`
		Fecha                  string  `json:"fecha"` // YYYY-MM-DD
	}

	var input rawControl
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	fecha, err := time.Parse("2006-01-02", input.Fecha)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Fecha inválida, usa YYYY-MM-DD"})
	}

	control := models.Control{
		PacienteID:             input.PacienteID,
		PesoKg:                 input.PesoKg,
		AlturaCm:               input.AlturaCm,
		IMC:                    input.IMC,
		PresionArterial:        input.PresionArterial,
		FrecuenciaCardiaca:     input.FrecuenciaCardiaca,
		FrecuenciaRespiratoria: input.FrecuenciaRespiratoria,
		TemperaturaC:           input.TemperaturaC,
		NivelGlucosa:           input.NivelGlucosa,
		SaturacionOxigeno:      input.SaturacionOxigeno,
		NotasGenerales:         input.NotasGenerales,
		Fecha:                  fecha,
	}

	if err := config.DB.Create(&control).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo crear el control"})
	}

	return c.Status(201).JSON(control)
}

func DeleteControl(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CN04")
	c.Locals("intCodeError", "F16")
	id := c.Params("id")
	if err := config.DB.Delete(&models.Control{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar control"})
	}
	return c.JSON(fiber.Map{"message": "Control eliminado"})
}

func GetControlUsuario(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CN05")
	c.Locals("intCodeError", "F34")
	id := c.Params("paciente_id")
	var controles []models.Control
	if err := config.DB.Find(&controles, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Control no encontrado"})
	}
	return c.JSON(controles)
}