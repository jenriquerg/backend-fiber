package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
	"github.com/jenriquerg/backend-fiber/middlewares"
)

func RecetaRoutes(app *fiber.App) {
	recetas := app.Group("/recetas")

	recetas.Get("/", middlewares.CheckPermission("get_recetas"), controllers.GetRecetas)
	recetas.Get("/:id", middlewares.CheckPermission("get_receta"), controllers.GetReceta)
	recetas.Post("/", middlewares.CheckPermission("add_receta"), controllers.CreateReceta)
	recetas.Put("/:id", middlewares.CheckPermission("update_receta"), controllers.UpdateReceta)
	recetas.Delete("/:id", middlewares.CheckPermission("delete_receta"), controllers.DeleteReceta)
}
