package models

// Like `Like` model is used to store like data,
// This model is used to store like data for both post and comment
// Like is bound to User and Post/Comment
type Like struct {
	ID     int  `json:"id"`
	UserID User `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	PostID Post `json:"post_id" gorm:"foreignKey:PostID;references:ID"`
}
