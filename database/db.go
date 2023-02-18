package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nevindra/sample-go-crud/models"
)

type DBStruct struct {
	Db *gorm.DB
}

var DB DBStruct

// ConnectDatabase connects to the database
func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=comic port=32768 sslmode=disable TimeZone=Asia/Singapore"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	log.Println("running migrations")

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Anime{}, &models.Review{})

	DB = DBStruct{
		Db: db,
	}
}
