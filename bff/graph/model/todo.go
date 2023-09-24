package model

type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	UserID *int64 `json:"userId"`
}
