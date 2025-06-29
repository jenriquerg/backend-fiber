package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
)

func ControlRoutes(app *fiber.App) {
	route := app.Group("/controles")

	route.Get("/", controllers.GetControles)
	route.Get("/:id", controllers.GetControl)
	route.Post("/", controllers.CreateControl)
	route.Delete("/:id", controllers.DeleteControl)
}