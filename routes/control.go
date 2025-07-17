package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/jenriquerg/backend-fiber/middlewares"
)

func ControlRoutes(app *fiber.App) {
	route := app.Group("/controles")

	route.Get("/", middlewares.CheckPermission("get_controles"), controllers.GetControles)
	route.Get("/:id", middlewares.CheckPermission("get_control"), controllers.GetControl)
	route.Post("/", middlewares.CheckPermission("add_control"), controllers.CreateControl)
	route.Delete("/:id", middlewares.CheckPermission("delete_control"), controllers.DeleteControl)
	route.Get("/paciente/:paciente_id", middlewares.CheckPermission("get_controles_paciente"), controllers.GetControlUsuario)
}
