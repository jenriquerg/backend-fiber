package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func CheckPermission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Token faltante"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{"error": "Formato de token inválido"})
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return getJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Token inválido"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "Token inválido"})
		}

		permissionsInterface, ok := claims["permisos"]
		if !ok {
			return c.Status(403).JSON(fiber.Map{"error": "Permisos no presentes en el token"})
		}

		permissionsRaw, ok := permissionsInterface.([]interface{})
		if !ok {
			return c.Status(403).JSON(fiber.Map{"error": "Permisos mal formateados"})
		}

		permissions := make([]string, 0, len(permissionsRaw))
		for _, p := range permissionsRaw {
			if str, ok := p.(string); ok {
				permissions = append(permissions, str)
			}
		}

		hasPermission := false
		for _, p := range permissions {
			if p == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return c.Status(403).JSON(fiber.Map{"error": "No tienes permiso para esta acción"})
		}

		c.Locals("userClaims", claims)

		return c.Next()
	}
}