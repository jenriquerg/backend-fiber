package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/jenriquerg/backend-fiber/middlewares"
)

func ExpedienteRoutes(app *fiber.App) {
	exp := app.Group("/expedientes")

	exp.Get("/", middlewares.CheckPermission("get_expedientes"), controllers.GetExpedientes)
	exp.Get("/:paciente_id", middlewares.CheckPermission("get_expediente"), controllers.GetExpediente)
	exp.Post("/", middlewares.CheckPermission("add_expediente"), controllers.CreateExpediente)
	exp.Put("/:id", middlewares.CheckPermission("update_expediente"), controllers.UpdateExpediente)
	exp.Delete("/:id", middlewares.CheckPermission("delete_expediente"), controllers.DeleteExpediente)
}
