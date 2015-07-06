package models

import (
	"database/sql"
	"log"

	"gopkg.in/gorp.v1"
)

type Todo struct {
	ID    int    `db:"todo_id"`
	Title string `db:"todo_title"`
	Done  bool   `db:"todo_done"`
}

func openClose(fn func(dbmap *gorp.DbMap)) {
	dbmap := initDb()
	defer dbmap.Db.Close()

	fn(dbmap)
}

func (t Todo) Insert() error {
	var err error
	openClose(func(dbmap *gorp.DbMap) {
		err = dbmap.Insert(&t)
	})
	return err
}

func TodoDone(id string) error {
	var err error
	var obj interface{}
	var count int64

	openClose(func(dbmap *gorp.DbMap) {
		obj, err = dbmap.Get(Todo{}, id)
		task := obj.(*Todo)
		task.Done = true

		count, err = dbmap.Update(task)
		log.Println("update count: %d", count)
	})

	return err
}

func AllTodos() ([]Todo, error) {
	var err error
	var todos []Todo
	// fetch all rows
	openClose(func(dbmap *gorp.DbMap) {
		_, err = dbmap.Select(&todos, "select * from todos order by todo_id")
	})

	return todos, err
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./app.db")
	CheckErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Todo{}, "todos").SetKeys(true, "ID")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}

// 初期データを投入する
func SetUpData() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	count, err := dbmap.SelectInt("select count (*) from todos")
	CheckErr(err, "select count (*) failed")

	if count == 0 {
		// delete any existing rows
		err = dbmap.TruncateTables()
		CheckErr(err, "Trucate failed")

		// insert todo
		t1 := &Todo{0, "buy milk", false}
		t2 := &Todo{1, "learn revel", true}

		err = dbmap.Insert(t1, t2)
		CheckErr(err, "Insert failed")
	}

	// delete row by PK
	// count, err := dbmap.Delete(&u1)
	// CheckErr(err, "Delete failed")
	// log.Println("Rows deleted:", count)

	// delete row manually via Exec
	// _, err = dbmap.Exec("delete from users where user_id=?", u2.Id)
	// CheckErr(err, "Exec failed")

	// confirm count is zero
	// count, err = dbmap.SelectInt("select count (*) from users")
	// CheckErr(err, "select count (*) failed")
	// log.Println("Row count - should be zero:", count)
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
