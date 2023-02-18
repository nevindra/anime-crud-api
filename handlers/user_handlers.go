package handlers

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
	"github.com/nevindra/sample-go-crud/database"
	"github.com/nevindra/sample-go-crud/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	database.DB.Db.Where("ID = ?", id).First(&user)
	// check if user is found
	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	// parse user input
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if user input email
	if user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email and password is required",
		})
	}

	// check email format with regex
	if !regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`).
		MatchString(user.Email) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid email format",
		})
	}

	// check if user email is already taken
	var tempUser models.User
	database.DB.Db.Where("email = ?", user.Email).First(&tempUser)
	if tempUser.Email != "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already taken",
		})
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	user.Password = string(hashedPassword)

	// generate nano id
	canonicID, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}

	user.ID = canonicID()

	// create user
	database.DB.Db.Create(&user)

	// return with 201 status code
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	// find user by string nano id
	id := c.Params("id")
	user := models.User{}

	// check user by id
	database.DB.Db.Where("ID = ?", id).First(&user)
	// check if user is found
	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// parse user input
	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if user input email
	if user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email and password is required",
		})
	}

	// check if user email is already taken
	var tempUser models.User
	database.DB.Db.Where("email = ?", user.Email).First(&tempUser)
	if tempUser.Email != "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already taken",
		})
	}

	// update the email
	database.DB.Db.Model(&user).Update("email", user.Email)
	return c.JSON(user)

}

func LoginUser(c *fiber.Ctx) error {
	// get user input
	user := new(models.User)

	// parse user input
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// check if user input email
	if user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email and password is required",
		})
	}

	// match email and password input with user in database
	var tempUser models.User
	database.DB.Db.Where("email = ?", user.Email).First(&tempUser)
	if tempUser.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(tempUser.Password), []byte(user.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// send user informatio in the response
	return c.JSON(fiber.Map{
		"user": tempUser,
	})
}
