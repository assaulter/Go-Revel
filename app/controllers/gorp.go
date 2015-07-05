package controllers

import (
	"database/sql"
	"log"

	"Go-Revel/app/models"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

func InitDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./app.db")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(models.Todo{}, "todos").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

// 初期データを投入する
func SetUpData() {
	dbmap := InitDb()
	defer dbmap.Db.Close()

	count, err := dbmap.SelectInt("select count (*) from todos")
	checkErr(err, "select count (*) failed")

	if count == 0 {
		// delete any existing rows
		err = dbmap.TruncateTables()
		checkErr(err, "Trucate failed")

		// insert todo
		t1 := &models.Todo{0, "buy milk", false}
		t2 := &models.Todo{1, "learn revel", true}

		err = dbmap.Insert(t1, t2)
		checkErr(err, "Insert failed")
	}

	// delete row by PK
	// count, err := dbmap.Delete(&u1)
	// checkErr(err, "Delete failed")
	// log.Println("Rows deleted:", count)

	// delete row manually via Exec
	// _, err = dbmap.Exec("delete from users where user_id=?", u2.Id)
	// checkErr(err, "Exec failed")

	// confirm count is zero
	// count, err = dbmap.SelectInt("select count (*) from users")
	// checkErr(err, "select count (*) failed")
	// log.Println("Row count - should be zero:", count)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
