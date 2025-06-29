package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
)

func RecetaRoutes(app *fiber.App) {
	recetas := app.Group("/recetas")

	recetas.Get("/", controllers.GetRecetas)
	recetas.Get("/:id", controllers.GetReceta)
	recetas.Post("/", controllers.CreateReceta)
	recetas.Put("/:id", controllers.UpdateReceta)
	recetas.Delete("/:id", controllers.DeleteReceta)
}
