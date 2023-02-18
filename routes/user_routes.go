package routes

import (
	handlers "github.com/nevindra/sample-go-crud/handlers"
)

// SetupUserRoutes is a function to set up all user routes and add it to `Routes` struct
func SetupUserRoutes(r *Routes) {
	r.Add("GET", "/users", handlers.GetUsers)
	r.Add("GET", "/users/:id", handlers.GetUser)
	r.Add("POST", "/users", handlers.CreateUser)
	r.Add("PUT", "/users/:id", handlers.UpdateUser)
	r.Add("POST", "/users/login", handlers.LoginUser)
}
