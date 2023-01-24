package models

type Video struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Likes       int64  `json:"likes"`
	Dislikes    int64  `json:"dislikes"`
	Views       int64  `json:"views"`
}
