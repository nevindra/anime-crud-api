package routes

import (
	"fmt"
	"github.com/nevindra/sample-go-crud/handlers"
)

// SetupAnimeRoutes is a function to set up all anime routes and add it to `Routes` struct
func SetupAnimeRoutes(r *Routes) {
	fmt.Println("Setting up anime routes")
	r.Add("GET", "/animes", handlers.GetAnimes)
	r.Add("GET", "/animes/:id", handlers.GetAnime)
	r.Add("POST", "/animes", handlers.CreateAnime)
	r.Add("PUT", "/animes/:id", handlers.UpdateAnime)
}
