package models

type Todo struct {
	Id    int    `db:"todo_id"`
	Title string `db:"todo_title"`
	Done  bool   `db:"todo_done"`
}
