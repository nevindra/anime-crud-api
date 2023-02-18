package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	Route []Route
}

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

// Add is a function to add a `route` to `Routes`
func (r *Routes) Add(method, path string, handler fiber.Handler) {
	r.Route = append(r.Route, Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
	fmt.Println("Added route: ", method, path)
}

// Register is a function to setup all handlers
func (r *Routes) Register(app *fiber.App) {
	for _, route := range r.Route {
		app.Add(route.Method, route.Path, route.Handler)
	}
	fmt.Println("Registered all handlers")
}

// Init create a function to initialize `Routes`
func Init(app *fiber.App) {
	routes := Routes{}
	SetupUserRoutes(&routes)
	SetupAnimeRoutes(&routes)
	SetupReviewRoutes(&routes)
	routes.Register(app)
}
