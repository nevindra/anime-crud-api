package models

// PostLike `PostLike` model is used to store like data,
// This model is used to store like data for both post and comment
// PostLike is bound to User and Post/Comment
type PostLike struct {
	ID     int  `json:"id"`
	UserID User `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	PostID Post `json:"post_id" gorm:"foreignKey:PostID;references:ID"`
}

type CommentLike struct {
	ID        int     `json:"id"`
	UserID    User    `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	CommentID Comment `json:"comment_id" gorm:"foreignKey:CommentID;references:ID"`
}
