package models

type Video struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Likes       int64  `json:"likes" db:"likes"`
	Dislikes    int64  `json:"dislikes" db:"dislikes"`
	Views       int64  `json:"views" db:"views"`
}
