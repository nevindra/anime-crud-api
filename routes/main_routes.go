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

// Setup is a function to setup all routes
func (r *Routes) Register(app *fiber.App) {
	for _, route := range r.Route {
		app.Add(route.Method, route.Path, route.Handler)
	}
	fmt.Println("Registered all routes")
}

// Create a function to add all routes from this folder to `Routes` struct
func (r *Routes) Setup(app *fiber.App) {
	// initialize user routes
	r.Add("GET", "/users", GetUsers)
	r.Add("GET", "/users/:id", GetUser)
	r.Add("POST", "/users", CreateUser)
	r.Add("PUT", "/users/:id", UpdateUser)

	// initialize `anime` routes
	r.Add("GET", "/animes", GetAnimes)
	r.Add("GET", "/animes/:id", GetAnime)
	r.Add("POST", "/animes", CreateAnime)
	r.Add("PUT", "/animes/:id", UpdateAnime)

	// initialize `review` routes
	r.Add("GET", "/reviews", GetReviews)
	r.Add("GET", "/reviews/:id", GetReview)
	r.Add("POST", "/reviews", CreateReview)
	r.Add("PUT", "/reviews/:id", UpdateReview)
}

// Create a function to initialize `Routes`
func Init(app *fiber.App) {
	routes := Routes{}
	routes.Setup(app)
	routes.Register(app)
}
