package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nevindra/sample-go-crud/database"
	"github.com/nevindra/sample-go-crud/models"
)

type LikeInput struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

// SendPostLike is a function to send a like to a post
// User ID and Post ID will be sent in the request body
func SendPostLike(c *fiber.Ctx) error {
	var like LikeInput
	// parse like input
	if err := c.BodyParser(&like); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if post is found
	post := models.Post{}
	database.DB.Db.Where("ID = ?", like.PostID).First(&post)
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// check if user is found
	user := models.User{}
	database.DB.Db.Where("ID = ?", like.UserID).First(&user)
	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// check if user is already liked the post
	likeModel := models.Like{}
	database.DB.Db.Where("user_id = ? AND post_id = ?", like.UserID, like.PostID).
		First(&likeModel)
	if likeModel.ID != 0 {
		// if user already liked the post, delete the like
		database.DB.Db.Delete(&likeModel)
		// decrease post like count
		database.DB.Db.Model(&post).Update("like_count", post.LikeCount-1)
		return c.JSON(fiber.Map{
			"message": "Like deleted",
		})
	} else {
		// if user not liked the post, create a new like
		AddLike := models.Like{
			UserID: user,
			PostID: post,
		}
		database.DB.Db.Create(&AddLike)
		// increase post like count
		database.DB.Db.Model(&post).Update("like_count", post.LikeCount+1)
		return c.JSON(fiber.Map{
			"message": "Like added",
		})
	}
}

// GetPostLikes is a function to get all likes of a post and all user who liked the post
func GetPostLikes(c *fiber.Ctx) error {
	// find post by string nano id
	id := c.Params("id")
	post := models.Post{}
	database.DB.Db.Where("ID = ?", id).First(&post)
	// check if post is found
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// get number of likes of the post
	likeCount := post.LikeCount

	// get all likes of the post
	var likes []models.Like
	database.DB.Db.Where("post_id = ?", id).Find(&likes)

	// get all users who liked the post
	var users []models.User
	for _, like := range likes {
		database.DB.Db.Where("ID = ?", like.UserID).First(&users)
	}

	// return all likes and users who liked the post
	return c.JSON(fiber.Map{
		"message":    "Post likes and users who liked the post",
		"like_count": likeCount,
		"users":      users,
	})
}

// GetCommentLikes is a function to get all likes of a comment and all user who liked the comment
func GetCommentLikes(c *fiber.Ctx) error {
	// find comment by string nano id
	id := c.Params("id")
	comment := models.Comment{}
	database.DB.Db.Where("ID = ?", id).First(&comment)
	// check if comment is found
	if comment.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Comment not found",
		})
	}

	// get number of likes of the comment
	likeCount := comment.LikeCount

	// return all likes and users who liked the comment
	return c.JSON(fiber.Map{
		"message":    "Comment likes",
		"like_count": likeCount,
	})
}
