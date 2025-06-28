package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/routes"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	routes.UsuarioRoutes(app)

	app.Listen(":3000")
}
