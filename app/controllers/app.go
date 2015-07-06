package controllers

import (
	m "Go-Revel/app/models"
	"log"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// show todos
func (c App) Index() revel.Result {
	todos, err := m.AllTodos()
	if err != nil {
		m.CheckErr(err, "AllTodos() has failed.")
	}
	log.Print(todos)

	return c.Render(todos)
}

// add todo
func (c App) Create(todo m.Todo) revel.Result {
	// Todo: Validation
	err := todo.Insert()
	if err != nil {
		m.CheckErr(err, "Insert() has failed.")
	}

	return c.Redirect(App.Index)
}

// change status to done
func (c App) Done(id string) revel.Result {
	log.Print(id)

	err := m.TodoDone(id)
	if err != nil {
		m.CheckErr(err, "TodoDone() has failed.")
	}

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
