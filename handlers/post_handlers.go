package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nevindra/sample-go-crud/database"
	"github.com/nevindra/sample-go-crud/models"
)

// GetPosts is a function to get all posts
func GetPosts(c *fiber.Ctx) error {
	posts := []models.Post{}
	database.DB.Db.Find(&posts)
	return c.JSON(posts)
}

// GetPost is a function to get a post by id
func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := models.Post{}
	database.DB.Db.Where("ID = ?", id).First(&post)
	// check if post is found
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	return c.JSON(post)
}

// CreatePost is a function to create a post
func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)

	// parse post input
	if err := c.BodyParser(post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if post input title
	if post.Title == "" || post.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Title and content is required",
		})
	}

	// create post
	database.DB.Db.Create(&post)

	// Send 201 status code and JSON response
	return c.Status(201).JSON(post)
}

// UpdatePost is a function to update a post by id
func UpdatePost(c *fiber.Ctx) error {
	// find post by string nano id
	id := c.Params("id")
	post := models.Post{}

	// get `User` model from context
	user := c.Locals("user").(models.User)

	// check post by id
	database.DB.Db.Where("ID = ?", id).First(&post)
	// check if post is found
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// check if the user is the owner of the post
	if post.UserID.ID != user.ID {
		return c.Status(403).JSON(fiber.Map{
			"message": "You are not allowed to update this post",
		})
	}

	// parse post input
	if err := c.BodyParser(&post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if post input title
	if post.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Content is required",
		})
	}

	// update the title
	database.DB.Db.Model(&post).Update("title", post.Title)
	return c.JSON(post)
}

// FindPostByUserID is a function to get all posts by user id
func FindPostByUserID(c *fiber.Ctx) error {
	// get user by params id
	id := c.Params("id")
	user := models.User{}
	database.DB.Db.Where("ID = ?", id).First(&user)

	// check if return is not empty
	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// get all posts by user id
	var posts []models.Post
	database.DB.Db.Where("user_id = ?", id).Find(&posts)

	// check if return is not empty
	if len(posts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// return posts
	return c.JSON(posts)
}

// SendLike is a function to send like to a post
func SendLike(c *fiber.Ctx) error {
	// get post by params id
	id := c.Params("id")
	post := models.Post{}
	database.DB.Db.Where("ID = ?", id).First(&post)

	// check if return is not empty
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// update like count
	database.DB.Db.Model(&post).Update("like_count", post.LikeCount+1)
	return c.JSON(post)
}

// SendDislike is a function to send dislike to a post
func SendDislike(c *fiber.Ctx) error {
	// get post by params id
	id := c.Params("id")
	post := models.Post{}
	database.DB.Db.Where("ID = ?", id).First(&post)

	// check if return is not empty
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// update dislike count
	database.DB.Db.Model(&post).Update("dislike_count", post.DislikeCount+1)
	return c.JSON(post)
}

// UndoLike is a function to undo like to a post
func UndoLike(c *fiber.Ctx) error {
	// get post by params id
	id := c.Params("id")
	post := models.Post{}
	database.DB.Db.Where("ID = ?", id).First(&post)

	// check if return is not empty
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// update like count
	database.DB.Db.Model(&post).Update("like_count", post.LikeCount-1)
	return c.JSON(post)
}

// UndoDislike is a function to undo dislike to a post
func UndoDislike(c *fiber.Ctx) error {
	// get post by params id
	id := c.Params("id")
	post := models.Post{}
	database.DB.Db.Where("ID = ?", id).First(&post)

	// check if return is not empty
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	// update dislike count
	database.DB.Db.Model(&post).Update("dislike_count", post.DislikeCount-1)
	return c.JSON(post)
}
