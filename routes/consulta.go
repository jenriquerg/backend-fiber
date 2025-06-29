package routes

import (
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConsultasRoutes(app *fiber.App) {
	consultas := app.Group("/consultas")

	consultas.Get("/", controllers.GetConsultas)
	consultas.Get("/:id", controllers.GetConsulta)
	consultas.Post("/", controllers.CreateConsulta)
	consultas.Put("/:id", controllers.UpdateConsulta)
	consultas.Delete("/:id", controllers.DeleteConsulta)
}