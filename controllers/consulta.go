package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
)

func GetConsultas(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "C01")
	c.Locals("intCodeError", "F03")
	var consultas []models.Consulta
	if err := config.DB.Find(&consultas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener consultas"})
	}
	return c.JSON(consultas)
}

func GetConsulta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "C02")
	c.Locals("intCodeError", "F04")
	id := c.Params("id")
	var consulta models.Consulta
	if err := config.DB.First(&consulta, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Consulta no encontrada"})
	}
	return c.JSON(consulta)
}

func CreateConsulta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "C03")
	c.Locals("intCodeError", "F05")
	type rawConsulta struct {
		IDConsultorio uint     `json:"id_consultorio"`
		IDMedico      uint     `json:"id_medico"`
		IDPaciente    uint     `json:"id_paciente"`
		Tipo          string   `json:"tipo"`
		Horario       string   `json:"horario"`
		Diagnostico   string   `json:"diagnostico"`
		Costo         *float64 `json:"costo"`
	}

	var input rawConsulta
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	fmt.Println("BODY:", string(c.Body()))

	horario, err := time.Parse(time.RFC3339, input.Horario)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha inválido. Usa RFC3339"})
	}

	consulta := models.Consulta{
		IDConsultorio: input.IDConsultorio,
		IDMedico:      input.IDMedico,
		IDPaciente:    input.IDPaciente,
		Tipo:          input.Tipo,
		Horario:       horario,
		Diagnostico:   input.Diagnostico,
		Costo:         input.Costo,
	}

	if err := config.DB.Create(&consulta).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al crear consulta"})
	}
	return c.Status(201).JSON(consulta)
}

func UpdateConsulta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "C04")
	c.Locals("intCodeError", "F06")
	id := c.Params("id")
	var consulta models.Consulta
	if err := config.DB.First(&consulta, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Consulta no encontrada"})
	}

	type rawConsulta struct {
		IDConsultorio uint     `json:"id_consultorio"`
		IDMedico      uint     `json:"id_medico"`
		IDPaciente    uint     `json:"id_paciente"`
		Tipo          string   `json:"tipo"`
		Horario       string   `json:"horario"`
		Diagnostico   string   `json:"diagnostico"`
		Costo         *float64 `json:"costo"`
	}

	var input rawConsulta
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	horario, err := time.Parse("2006-01-02T15:04:05", input.Horario)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha inválido. Usa YYYY-MM-DDTHH:MM:SS"})
	}

	consulta.IDConsultorio = input.IDConsultorio
	consulta.IDMedico = input.IDMedico
	consulta.IDPaciente = input.IDPaciente
	consulta.Tipo = input.Tipo
	consulta.Horario = horario
	consulta.Diagnostico = input.Diagnostico
	consulta.Costo = input.Costo

	if err := config.DB.Save(&consulta).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo actualizar la consulta"})
	}
	return c.JSON(consulta)
}

func DeleteConsulta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "C05")
	c.Locals("intCodeError", "F07")
	id := c.Params("id")
	if err := config.DB.Delete(&models.Consulta{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar consulta"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Consulta eliminada"})
}

func GetConsultasByPaciente(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "C06")
	c.Locals("intCodeError", "F32")
	id := c.Params("id")

	var consultas []models.Consulta
	if err := config.DB.Where("id_paciente = ?", id).Find(&consultas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener las consultas del paciente"})
	}

	return c.JSON(consultas)
}

func GetConsultasByMedico(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "C07")
	c.Locals("intCodeError", "F33")
	id := c.Params("id")

	var consultas []models.Consulta
	if err := config.DB.Where("id_medico = ?", id).Find(&consultas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener las consultas del médico"})
	}

	return c.JSON(consultas)
}
