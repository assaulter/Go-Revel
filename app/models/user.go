package models

type User struct {
	Id   int    `db:"user_id"`
	Name string `db:user_name,size:24`
}
