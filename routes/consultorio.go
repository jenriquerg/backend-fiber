package routes

import (
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConsultorioRoutes(app *fiber.App) {
	consultorios := app.Group("/consultorios")

	consultorios.Get("/", controllers.GetConsultorios)
	consultorios.Get("/:id", controllers.GetConsultorio)
	consultorios.Post("/", controllers.CreateConsultorio)
	consultorios.Put("/:id", controllers.UpdateConsultorio)
	consultorios.Delete("/:id", controllers.DeleteConsultorio)
}