package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
)

func GetConsultorios(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CT01")
	c.Locals("intCodeError", "F08")
	var consultorios []models.Consultorio
	if err := config.DB.Find(&consultorios).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener consultorios"})
	}
	return c.JSON(consultorios)
}

func GetConsultorio(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CT02")
	c.Locals("intCodeError", "F09")	
	id := c.Params("id")
	var consultorio models.Consultorio
	if err := config.DB.First(&consultorio, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Consultorio no encontrado"})
	}
	return c.JSON(consultorio)
}

func CreateConsultorio(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CT03")
	c.Locals("intCodeError", "F10")
	var input models.Consultorio
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if input.Status == "" {
		input.Status = "disponible"
	}

	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al crear consultorio"})
	}
	return c.Status(201).JSON(input)
}

func UpdateConsultorio(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CT04")
	c.Locals("intCodeError", "F11")
	id := c.Params("id")
	var consultorio models.Consultorio
	if err := config.DB.First(&consultorio, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Consultorio no encontrado"})
	}

	var input models.Consultorio
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	consultorio.Nombre = input.Nombre
	consultorio.Tipo = input.Tipo
	consultorio.IDMedico = input.IDMedico
	consultorio.Status = input.Status
	consultorio.Ubicacion = input.Ubicacion
	consultorio.Horario = input.Horario

	if err := config.DB.Save(&consultorio).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo actualizar el consultorio"})
	}
	return c.JSON(consultorio)
}

func DeleteConsultorio(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "CT05")
	c.Locals("intCodeError", "F12")
	id := c.Params("id")
	if err := config.DB.Delete(&models.Consultorio{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar consultorio"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Consultorio eliminado"})
}