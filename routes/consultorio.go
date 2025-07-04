package routes

import (
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/middlewares"
)

func ConsultorioRoutes(app *fiber.App) {
	consultorios := app.Group("/consultorios")

	consultorios.Get("/", middlewares.CheckPermission("get_consultorios"), controllers.GetConsultorios)
	consultorios.Get("/:id", middlewares.CheckPermission("get_consultorio"), controllers.GetConsultorio)
	consultorios.Post("/", middlewares.CheckPermission("add_consultorio"), controllers.CreateConsultorio)
	consultorios.Put("/:id", middlewares.CheckPermission("update_consultorio"), controllers.UpdateConsultorio)
	consultorios.Delete("/:id", middlewares.CheckPermission("delete_consultorio"), controllers.DeleteConsultorio)
}
