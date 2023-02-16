package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nevindra/sample-go-crud/database"
	"github.com/nevindra/sample-go-crud/models"
	"gorm.io/gorm/logger"
)

func GetAnimes(c *fiber.Ctx) error {
	animes := []models.Anime{}
	database.DB.Db.Find(&animes)
	return c.JSON(animes)
}

func GetAnime(c *fiber.Ctx) error {
	id := c.Params("id")
	anime := models.Anime{}
	database.DB.Db.Find(&anime, id)
	// check if anime is found
	if anime.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Anime not found",
		})
	}

	return c.JSON(anime)
}

func CreateAnime(c *fiber.Ctx) error {
	anime := new(models.Anime)

	//parse inpt
	if err := c.BodyParser(anime); err != nil {
		// log error
		logger.Error(err)

		// return 503 with body error
		return c.Status(503).SendString(err.Error())
	}

	// check if every field is filled
	if anime.Title == "" || anime.Genre == "" || anime.Episode == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Every field is required",
		})
	}

	// check if anime is already exist
	var tempAnime models.Anime
	database.DB.Db.Where("title = ?", anime.Title).First(&tempAnime)
	if tempAnime.Title != "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Anime is already exist",
		})
	}

	// create anime
	database.DB.Db.Create(&anime)

	return c.JSON(anime)
}

func UpdateAnime(c *fiber.Ctx) error {
	id := c.Params("id")
	anime := models.Anime{}
	database.DB.Db.Find(&anime, id)
	// check if anime is found
	if anime.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Anime not found",
		})
	}

	// parse input
	if err := c.BodyParser(&anime); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if every field is filled
	if anime.Title == "" || anime.Genre == "" || anime.Episode == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Every field is required",
		})
	}

	// check if anime is already exist
	var tempAnime models.Anime
	database.DB.Db.Where("title = ?", anime.Title).First(&tempAnime)
	if tempAnime.Title != "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Anime is already exist",
		})
	}

	// update anime
	database.DB.Db.Model(&anime).Updates(anime)

	return c.JSON(anime)
}
