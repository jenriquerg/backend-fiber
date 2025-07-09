package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/models"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Ejecutar controlador
		err := c.Next()

		// Duración de la petición
		duration := time.Since(start).Milliseconds()

		// Obtener info de la petición
		method := c.Method()
		path := c.OriginalURL()
		status := c.Response().StatusCode()
		ip := c.IP()

		// ID del usuario (desde token si existe)
		var userID string
		if claims, ok := c.Locals("userClaims").(map[string]interface{}); ok {
			if uid, ok := claims["sub"].(string); ok {
				userID = uid
			}
		}

		// IntCode personalizado (puede venir como string, por eso usamos interface{})
		var intCode string
		if status >= 200 && status < 300 {
			if code, ok := c.Locals("intCodeSuccess").(string); ok {
				intCode = code
			}
		} else {
			if code, ok := c.Locals("intCodeError").(string); ok {
				intCode = code
			}
		}

		// Crear registro de log
		log := models.RequestLog{
			Method:     method,
			Path:       path,
			StatusCode: status,
			DurationMs: duration,
			IPAddress:  ip,
			UserID:     userID,
			IntCode:    intCode,
			CreatedAt:  time.Now(),
		}

		// Guardar en la base de datos
		config.DB.Create(&log)

		return err
	}
}