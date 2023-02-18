package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/nevindra/sample-go-crud/database"
	"github.com/nevindra/sample-go-crud/routes"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDatabase()
	app := fiber.New()

	// setup routes
	routes.Init(app)

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
