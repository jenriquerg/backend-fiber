// controllers/auth_controller.go
package controllers

import (
	"os"
	"time"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/pquerna/otp/totp"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
)

func Register(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "R01")
	c.Locals("intCodeError", "F01")
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

	// Validar contraseña segura
	if len(input.Password) < 12 {
		return c.Status(400).JSON(fiber.Map{"error": "La contraseña debe tener al menos 12 caracteres"})
	}

	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	hasSymbol := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString

	if !hasLetter(input.Password) {
		return c.Status(400).JSON(fiber.Map{"error": "La contraseña debe contener al menos una letra"})
	}
	if !hasNumber(input.Password) {
		return c.Status(400).JSON(fiber.Map{"error": "La contraseña debe contener al menos un número"})
	}
	if !hasSymbol(input.Password) {
		return c.Status(400).JSON(fiber.Map{"error": "La contraseña debe contener al menos un símbolo"})
	}

	fecha, err := time.Parse("2006-01-02", input.FechaNacimiento)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Fecha inválida (usa YYYY-MM-DD)"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al encriptar contraseña"})
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "backend-fiber",
		AccountName: input.Correo,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error generando clave TOTP"})
	}

	usuario := models.Usuario{
		Nombre:          input.Nombre,
		Apellidos:       input.Apellidos,
		Tipo:            "paciente",
		FechaNacimiento: fecha,
		Genero:          input.Genero,
		Correo:          input.Correo,
		Password:        string(hash),
		SecretTOTP:      key.Secret(),
		CreadoEn:        time.Now(),
		ActualizadoEn:   time.Now(),
	}

	if err := config.DB.Create(&usuario).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo registrar el usuario"})
	}

	qrURL := key.URL()

	return c.Status(201).JSON(fiber.Map{
		"message": "Registro exitoso",
		"qr_url":  qrURL,
	})
}


func Login(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "L01")
	c.Locals("intCodeError", "F02")
	var input struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
		TOTPCode string `json:"totp_code"`
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

	// Validar código TOTP
	if !totp.Validate(input.TOTPCode, usuario.SecretTOTP) {
		return c.Status(401).JSON(fiber.Map{"error": "Código TOTP inválido"})
	}

	// Obtener permisos desde la base de datos usando el tipo
	var rolID int
	err := config.DB.
		Raw("SELECT id FROM roles WHERE nombre = ?", usuario.Tipo).
		Scan(&rolID).Error
	if err != nil || rolID == 0 {
		return c.Status(500).JSON(fiber.Map{"error": "Rol no encontrado"})
	}

	// Obtener lista de permisos (nombres)
	var permisos []string
	err = config.DB.
		Raw(`
			SELECT p.nombre
			FROM permisos p
			JOIN rol_permisos rp ON rp.permiso_id = p.id
			WHERE rp.rol_id = ?
		`, rolID).
		Scan(&permisos).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener permisos"})
	}

	// Crear Access Token (20 minutos)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       usuario.ID,
		"correo":   usuario.Correo,
		"tipo":     usuario.Tipo,
		"permisos": permisos,
		"exp":      time.Now().Add(20 * time.Minute).Unix(),
	})
	accessTokenStr, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al generar access token"})
	}

	// Crear Refresh Token (40 minutos)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     usuario.ID,
		"correo": usuario.Correo,
		"type":   "refresh",
		"exp":    time.Now().Add(40 * time.Minute).Unix(),
	})
	refreshTokenStr, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al generar refresh token"})
	}

	return c.JSON(fiber.Map{
		"access_token":  accessTokenStr,
		"refresh_token": refreshTokenStr,
	})
}
