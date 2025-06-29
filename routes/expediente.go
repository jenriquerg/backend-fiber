package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
)

func ExpedienteRoutes(app *fiber.App) {
	r := app.Group("/expedientes")

	r.Get("/", controllers.GetExpedientes)
	r.Get("/:id", controllers.GetExpediente)
	r.Post("/", controllers.CreateExpediente)
	r.Put("/:id", controllers.UpdateExpediente)
	r.Delete("/:id", controllers.DeleteExpediente)
}
