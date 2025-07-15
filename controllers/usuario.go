package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetUsuarios(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "U01")
	c.Locals("intCodeError", "F27")
	var usuarios []models.Usuario
	if err := config.DB.Find(&usuarios).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener usuarios"})
	}
	return c.JSON(usuarios)
}


func GetUsuario(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "U02")
	c.Locals("intCodeError", "F28")
	id := c.Params("id")
	var usuario models.Usuario
	if err := config.DB.First(&usuario, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}
	return c.JSON(usuario)
}

func CreateUsuario(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "U03")
	c.Locals("intCodeError", "F29")
	type rawUsuario struct {
		Nombre          string `json:"nombre"`
		Apellidos       string `json:"apellidos"`
		Tipo            string `json:"tipo"`
		FechaNacimiento string `json:"fecha_nacimiento"`
		Genero          string `json:"genero"`
		Correo          string `json:"correo"`
		Password        string `json:"password"`
	}

	var input rawUsuario

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	fecha, err := time.Parse("2006-01-02", input.FechaNacimiento)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha inválido. Usa YYYY-MM-DD"})
	}

	// Hashear contraseña
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al encriptar contraseña"})
	}

	// Generar clave secreta TOTP
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MiApp", // Cambiar por tu app
		AccountName: input.Correo,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error generando clave TOTP"})
	}

	usuario := models.Usuario{
		Nombre:          input.Nombre,
		Apellidos:       input.Apellidos,
		Tipo:            input.Tipo,
		FechaNacimiento: fecha,
		Genero:          input.Genero,
		Correo:          input.Correo,
		Password:        string(hashed),
		SecretTOTP:      key.Secret(),
		CreadoEn:        time.Now(),
		ActualizadoEn:   time.Now(),
	}

	if err := config.DB.Create(&usuario).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo crear el usuario"})
	}

	qrURL := key.URL()

	response := struct {
		ID              uint      `json:"id"`
		Nombre          string    `json:"nombre"`
		Apellidos       string    `json:"apellidos"`
		Tipo            string    `json:"tipo"`
		FechaNacimiento string    `json:"fecha_nacimiento"`
		Genero          string    `json:"genero"`
		Correo          string    `json:"correo"`
		CreadoEn        time.Time `json:"creado_en"`
		ActualizadoEn   time.Time `json:"actualizado_en"`
		QRURL           string    `json:"qr_url"`
		Message         string    `json:"message"`
	}{
		ID:              usuario.ID,
		Nombre:          usuario.Nombre,
		Apellidos:       usuario.Apellidos,
		Tipo:            usuario.Tipo,
		FechaNacimiento: usuario.FechaNacimiento.Format("2006-01-02"),
		Genero:          usuario.Genero,
		Correo:          usuario.Correo,
		CreadoEn:        usuario.CreadoEn,
		ActualizadoEn:   usuario.ActualizadoEn,
		QRURL:           qrURL,
		Message:         "Usuario creado exitosamente. Escanea el QR para configurar tu autenticador.",
	}

	return c.Status(201).JSON(response)

}

func UpdateUsuario(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "U04")
	c.Locals("intCodeError", "F30")
	id := c.Params("id")

	type rawUsuario struct {
		Nombre          string `json:"nombre"`
		Apellidos       string `json:"apellidos"`
		Tipo            string `json:"tipo"`
		FechaNacimiento string `json:"fecha_nacimiento"`
		Genero          string `json:"genero"`
		Correo          string `json:"correo"`
		Password        string `json:"password"`
	}

	var input rawUsuario

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	fecha, err := time.Parse("2006-01-02", input.FechaNacimiento)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha inválido. Usa YYYY-MM-DD"})
	}

	usuario := new(models.Usuario)

	// Buscar usuario existente
	if err := config.DB.First(&usuario, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}

	// Hashear contraseña solo si viene en el body y no está vacía
	if input.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error al encriptar contraseña"})
		}
		usuario.Password = string(hash)
	}

	// Actualizar campos (sin tocar SecretTOTP)
	usuario.Nombre = input.Nombre
	usuario.Apellidos = input.Apellidos
	usuario.Tipo = input.Tipo
	usuario.FechaNacimiento = fecha
	usuario.Genero = input.Genero
	usuario.Correo = input.Correo
	usuario.ActualizadoEn = time.Now()

	// Guardar cambios
	if err := config.DB.Save(&usuario).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo actualizar el usuario"})
	}

	// Retornamos usuario filtrado sin password ni secret (igual que en Create)
	response := struct {
		ID              uint      `json:"id"`
		Nombre          string    `json:"nombre"`
		Apellidos       string    `json:"apellidos"`
		Tipo            string    `json:"tipo"`
		FechaNacimiento string    `json:"fecha_nacimiento"`
		Genero          string    `json:"genero"`
		Correo          string    `json:"correo"`
		CreadoEn        time.Time `json:"creado_en"`
		ActualizadoEn   time.Time `json:"actualizado_en"`
	}{
		ID:              usuario.ID,
		Nombre:          usuario.Nombre,
		Apellidos:       usuario.Apellidos,
		Tipo:            usuario.Tipo,
		FechaNacimiento: usuario.FechaNacimiento.Format("2006-01-02"),
		Genero:          usuario.Genero,
		Correo:          usuario.Correo,
		CreadoEn:        usuario.CreadoEn,
		ActualizadoEn:   usuario.ActualizadoEn,
	}

	return c.JSON(response)
}

func DeleteUsuario(c *fiber.Ctx) error {
	c.Locals("intCodeSuccess", "U05")
	c.Locals("intCodeError", "F31")
	id := c.Params("id")

	if err := config.DB.Delete(&models.Usuario{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar usuario"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Usuario eliminado"})
}
