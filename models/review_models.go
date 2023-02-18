package models

// Review model is used to store review data,
// User ID and Anime ID is foreign key from `User` and `Anime` model
type Review struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	UserID  User   `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	AnimeID Anime  `json:"anime_id" gorm:"foreignKey:AnimeID;references:ID"`
	Rating  int    `json:"rating"`
}
