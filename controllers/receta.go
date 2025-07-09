package controllers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
)

func GetRecetas(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "R01")
	c.Locals("intCodeError", "F22")
	var recetas []models.Receta
	if err := config.DB.Find(&recetas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener recetas"})
	}
	return c.JSON(recetas)
}

func GetReceta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "R02")
	c.Locals("intCodeError", "F23")
	id := c.Params("id")
	var receta models.Receta
	if err := config.DB.First(&receta, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Receta no encontrada"})
	}
	return c.JSON(receta)
}

func CreateReceta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "R03")
	c.Locals("intCodeError", "F24")
	type rawReceta struct {
		IdConsulta  uint   `json:"id_consulta"`
		Fecha       string `json:"fecha"` // Formato: YYYY-MM-DD
		IdMedico    uint   `json:"id_medico"`
		Medicamento string `json:"medicamento"`
		Dosis       string `json:"dosis"`
	}

	var input rawReceta
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	fecha, err := time.Parse("2006-01-02", input.Fecha)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha inválido. Usa YYYY-MM-DD"})
	}

	receta := models.Receta{
		IdConsulta:  input.IdConsulta,
		Fecha:       fecha,
		IdMedico:    input.IdMedico,
		Medicamento: input.Medicamento,
		Dosis:       input.Dosis,
	}

	if err := config.DB.Create(&receta).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo crear la receta"})
	}

	return c.Status(201).JSON(receta)
}

func UpdateReceta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "R04")
	c.Locals("intCodeError", "F25")
	id := c.Params("id")

	var input struct {
		IdConsulta  uint   `json:"id_consulta"`
		Fecha       string `json:"fecha"`
		IdMedico    uint   `json:"id_medico"`
		Medicamento string `json:"medicamento"`
		Dosis       string `json:"dosis"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	fecha, err := time.Parse("2006-01-02", input.Fecha)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha inválido"})
	}

	var receta models.Receta
	if err := config.DB.First(&receta, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Receta no encontrada"})
	}

	receta.IdConsulta = input.IdConsulta
	receta.Fecha = fecha
	receta.IdMedico = input.IdMedico
	receta.Medicamento = input.Medicamento
	receta.Dosis = input.Dosis

	if err := config.DB.Save(&receta).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar receta"})
	}

	return c.JSON(receta)
}

func DeleteReceta(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "R05")
	c.Locals("intCodeError", "F26")
	id := c.Params("id")
	if err := config.DB.Delete(&models.Receta{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar receta"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Receta eliminada"})
}
