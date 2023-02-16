package models

type Anime struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Episodes    int    `json:"episodes"`
	Rating      int    `json:"rating"`
}
