// controllers/auth_controller.go
package controllers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
)

func Register(c *fiber.Ctx) error {
	var input struct {
		Nombre          string `json:"nombre"`
		Apellidos       string `json:"apellidos"`
		FechaNacimiento string `json:"fecha_nacimiento"`
		Genero          string `json:"genero"`
		Correo          string `json:"correo"`
		Password        string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	fecha, err := time.Parse("2006-01-02", input.FechaNacimiento)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Fecha inválida (usa YYYY-MM-DD)"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al encriptar contraseña"})
	}

	usuario := models.Usuario{
		Nombre:          input.Nombre,
		Apellidos:       input.Apellidos,
		Tipo:            "paciente",
		FechaNacimiento: fecha,
		Genero:          input.Genero,
		Correo:          input.Correo,
		Password:        string(hash),
		CreadoEn:        time.Now(),
		ActualizadoEn:   time.Now(),
	}

	if err := config.DB.Create(&usuario).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo registrar el usuario"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Registro exitoso"})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	var usuario models.Usuario
	if err := config.DB.Where("correo = ?", input.Correo).First(&usuario).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Contraseña incorrecta"})
	}

	// Generar token JWT
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    usuario.ID,
		"email": usuario.Correo,
		"tipo":  usuario.Tipo,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // 3 días
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al generar token"})
	}

	return c.JSON(fiber.Map{"token": tokenStr})
}