package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/controllers"
)

func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}
