package routes

import (
	"fmt"
	"github.com/nevindra/sample-go-crud/handlers"
)

// SetupReviewRoutes is a function to set up all review routes and add it to `Routes` struct
func SetupReviewRoutes(r *Routes) {
	fmt.Println("Setting up review routes")
	r.Add("GET", "/reviews", handlers.GetReviews)
	r.Add("GET", "/reviews/:id", handlers.GetReview)
	r.Add("POST", "/reviews", handlers.CreateReview)
	r.Add("PUT", "/reviews/:id", handlers.UpdateReview)
}
