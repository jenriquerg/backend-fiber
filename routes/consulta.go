package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/jenriquerg/backend-fiber/middlewares"
)

func ConsultasRoutes(app *fiber.App) {
	consultas := app.Group("/consultas")

	consultas.Get("/", middlewares.CheckPermission("get_consultas"), controllers.GetConsultas)
	consultas.Get("/:id", middlewares.CheckPermission("get_consulta"), controllers.GetConsulta)
	consultas.Post("/", middlewares.CheckPermission("add_consulta"), controllers.CreateConsulta)
	consultas.Put("/:id", middlewares.CheckPermission("update_consulta"), controllers.UpdateConsulta)
	consultas.Delete("/:id", middlewares.CheckPermission("delete_consulta"), controllers.DeleteConsulta)
}
