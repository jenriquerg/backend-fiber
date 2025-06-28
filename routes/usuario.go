package routes

import (
	"github.com/jenriquerg/backend-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func UsuarioRoutes(app *fiber.App) {
	usuarios := app.Group("/usuarios")

	usuarios.Get("/", controllers.GetUsuarios)
	usuarios.Get("/:id", controllers.GetUsuario)
	usuarios.Post("/", controllers.CreateUsuario)
	usuarios.Put("/:id", controllers.UpdateUsuario)
	usuarios.Delete("/:id", controllers.DeleteUsuario)
}