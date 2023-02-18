package routes

import (
	"fmt"
	"github.com/nevindra/sample-go-crud/handlers"

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

// Setup create a function to add all handlers with `handlers`
func (r *Routes) Setup() {
	// initialize user handlers
	r.Add("GET", "/users", handlers.GetUsers)
	r.Add("GET", "/users/:id", handlers.GetUser)
	r.Add("POST", "/users", handlers.CreateUser)
	r.Add("PUT", "/users/:id", handlers.UpdateUser)

	// initialize `anime` handlers
	r.Add("GET", "/animes", handlers.GetAnimes)
	r.Add("GET", "/animes/:id", handlers.GetAnime)
	r.Add("POST", "/animes", handlers.CreateAnime)
	r.Add("PUT", "/animes/:id", handlers.UpdateAnime)

	// initialize `review` handlers
	r.Add("GET", "/reviews", handlers.GetReviews)
	r.Add("GET", "/reviews/:id", handlers.GetReview)
	r.Add("POST", "/reviews", handlers.CreateReview)
	r.Add("PUT", "/reviews/:id", handlers.UpdateReview)
}

// Init create a function to initialize `Routes`
func Init(app *fiber.App) {
	routes := Routes{}
	routes.Setup()
	routes.Register(app)
}
