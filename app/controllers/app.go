package controllers

import (
	"Go-Revel/app/models"
	"log"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// show todos
func (c App) Index() revel.Result {
	// initialize the DbMap
	dbmap := InitDb()
	defer dbmap.Db.Close()
	// set up initial data
	// SetUpData()

	// fetch all rows
	var todos []models.Todo
	var err error
	_, err = dbmap.Select(&todos, "select * from todos order by todo_id")
	checkErr(err, "Select failed")
	log.Printf("All rows:")
	for i, u := range todos {
		log.Printf(" %d: %v\n", i, u)
	}

	return c.Render(todos)
}

// add todo
func (c App) Create(todo models.Todo) revel.Result {
	// Todo: Validation
	// initialize the DbMap
	dbmap := InitDb()
	defer dbmap.Db.Close()

	var err error
	err = dbmap.Insert(&todo)
	checkErr(err, "Insert failed")

	return c.Redirect(App.Index)
}

// change status to done
func (c App) Done(id string) revel.Result {
	log.Print(id)

	dbmap := InitDb()
	defer dbmap.Db.Close()

	obj, err := dbmap.Get(models.Todo{}, id)
	checkErr(err, "Get failed")
	t := obj.(*models.Todo)
	t.Done = true

	count, err := dbmap.Update(t)
	checkErr(err, "Update failed")
	log.Print(count)

	return c.Redirect(App.Index)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}
