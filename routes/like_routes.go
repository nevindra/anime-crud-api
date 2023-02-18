package routes

import (
	"fmt"
	handler "github.com/nevindra/sample-go-crud/handlers"
)

// LikeRoutes is a function to setup all like routes
func LikeRoutes(r *Routes) {
	fmt.Println("Setting up like routes")
	r.Add("GET", "/like/comments/:id", handler.GetCommentLikes)
	r.Add("GET", "/like/posts/:id", handler.GetPostLikes)
	r.Add("POST", "/like", handler.SendPostLike)
	r.Add("POST", "/like", handler.SendCommentLike)
}
