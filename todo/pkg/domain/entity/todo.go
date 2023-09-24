package entity

type Todo struct {
	ID     int64  `db:"id"`
	Title  string `db:"title"`
	Done   bool   `db:"done"`
	UserID int64  `db:"user_id"`
}
