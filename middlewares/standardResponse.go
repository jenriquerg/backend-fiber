package middlewares

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func StandardResponse() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return err
		}

		status := c.Response().StatusCode()
		bodyBytes := c.Response().Body()

		var parsed interface{}
		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
				parsed = string(bodyBytes)
			}
		}

		// Transformar a arreglo siempre
		var data []interface{}
		switch v := parsed.(type) {
		case nil:
			data = []interface{}{}
		case []interface{}:
			data = v
		default:
			data = []interface{}{v}
		}

		// Obtener intCode (string o int)
		var intCode interface{}
		if status >= 200 && status < 300 {
			if code := c.Locals("intCodeSuccess"); code != nil {
				switch v := code.(type) {
				case int, string:
					intCode = v
				default:
					intCode = 1000 + status
				}
			} else {
				intCode = 1000 + status
			}
		} else {
			if code := c.Locals("intCodeError"); code != nil {
				switch v := code.(type) {
				case int, string:
					intCode = v
				default:
					intCode = 4000 + status
				}
			} else {
				switch {
				case status >= 400 && status < 500:
					intCode = 4000 + status
				case status >= 500 && status < 600:
					intCode = 5000 + status
				default:
					intCode = 9000 + status
				}
			}
		}

		// Empaquetar respuesta final
		wrapped, _ := json.Marshal(fiber.Map{
			"statusCode": status,
			"intCode":    intCode,
			"data":       data,
		})
		c.Response().SetBody(wrapped)
		return nil
	}
}
