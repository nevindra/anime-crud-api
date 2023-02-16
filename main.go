package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/nevindra/sample-go-crud/database"
	"github.com/nevindra/sample-go-crud/routes"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

// setup user routes
func setupUserRoutes(app *fiber.App) {
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Post("/api/users", routes.CreateUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Post("/api/login", routes.Login)
}

func main() {
	database.ConnectDatabase()
	app := fiber.New()
	setupUserRoutes(app)

	// enable cors
	app.Use(cors.New())

	// rate limit
	app.Use(limiter.New(limiter.Config{
		Max: 2,
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// start server
	app.Listen(":3000")
}
