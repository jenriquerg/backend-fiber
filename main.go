package main

import (
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jenriquerg/backend-fiber/config"
	"github.com/jenriquerg/backend-fiber/routes"
	"github.com/jenriquerg/backend-fiber/middlewares"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	config.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
    // AllowOrigins: "http://localhost:4200",
    AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
}))

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	}))

	app.Use(middlewares.RequestLogger())
	app.Use(middlewares.StandardResponse())

	routes.UsuarioRoutes(app)
	routes.AuthRoutes(app)
	routes.ConsultorioRoutes(app)
	routes.ConsultasRoutes(app)
	routes.RecetaRoutes(app)
	routes.ExpedienteRoutes(app)
	routes.ControlRoutes(app)

	app.Listen(":3000")
}