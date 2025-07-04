package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/jenriquerg/backend-fiber/middlewares"
)

func UsuarioRoutes(app *fiber.App) {
	usuarios := app.Group("/usuarios")

	usuarios.Get("/", middlewares.CheckPermission("get_users"), controllers.GetUsuarios)
	usuarios.Get("/:id", middlewares.CheckPermission("get_users"), controllers.GetUsuario)
	usuarios.Post("/", middlewares.CheckPermission("add_user"), controllers.CreateUsuario)
	usuarios.Put("/:id", middlewares.CheckPermission("update_user"), controllers.UpdateUsuario)
	usuarios.Delete("/:id", middlewares.CheckPermission("delete_user"), controllers.DeleteUsuario)
}
