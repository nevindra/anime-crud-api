package models

// Post `Post` model is used to store post data,
// User ID is foreign key from `User` model
// User is able to create many posts
// User is able to like and dislike many posts
type Post struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	UserID       User   `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	LikeCount    int    `json:"like_count"`
	DislikeCount int    `json:"dislike"`
}

// Comment `Comment` model is used to store comment data,
// User ID and Post ID is foreign key from `User` and `Post` model
// User is able to create many comments
// User is able to like and dislike many comments
type Comment struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	UserID       User   `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	PostID       Post   `json:"post_id" gorm:"foreignKey:PostID;references:ID"`
	LikeCount    int    `json:"like_count"`
	DislikeCount int    `json:"dislike"`
}
