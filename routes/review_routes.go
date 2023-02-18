package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nevindra/sample-go-crud/database"
	"github.com/nevindra/sample-go-crud/models"
)

// GetReviews get all reviews
func GetReviews(c *fiber.Ctx) error {
	reviews := []models.Review{}
	database.DB.Db.Find(&reviews)
	return c.JSON(reviews)
}

// GetReview get review by id
func GetReview(c *fiber.Ctx) error {
	id := c.Params("id")
	review := models.Review{}
	// find review by id
	database.DB.Db.Where("id = ?", id).First(&review)
	// check if review is found
	if review.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Review not found",
		})
	}

	return c.JSON(review)
}

func CreateReview(c *fiber.Ctx) error {
	review := new(models.Review)

	//parse inpt
	if err := c.BodyParser(review); err != nil {
		// return 503 with body error
		return c.Status(503).SendString(err.Error())
	}

	// check if user input anime id is empty
	if review.AnimeID.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Anime ID is required",
		})
	}

	// check if user input rating is between 1 and 5
	if review.Rating < 1 || review.Rating > 5 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Rating must be between 1 and 5",
		})
	}

	// check if user input review is empty
	if review.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Review is required",
		})
	}

	// create review
	database.DB.Db.Create(&review)

	// return with status 201 and review
	return c.Status(201).JSON(review)
}

// create a function to update the review
func UpdateReview(c *fiber.Ctx) error {
	// find review by string id
	id := c.Params("id")
	review := models.Review{}

	// check review by id
	database.DB.Db.Where("id = ?", id).First(&review)
	// check if review is found
	if review.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Review not found",
		})
	}

	// parse user input
	if err := c.BodyParser(&review); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if user input rating is between 1 and 5
	if review.Rating < 1 || review.Rating > 5 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Rating must be between 1 and 5",
		})
	}

	// check if user input review is empty
	if review.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Review is required",
		})
	}

	// update the review
	database.DB.Db.Model(&review).Updates(review)
	return c.JSON(review)
}
