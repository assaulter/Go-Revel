package controllers

import (
	"Go-Revel/app/models"

	"log"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// initialize the DbMap
	dbmap := InitDb()
	defer dbmap.Db.Close()

	// delete any existing rows
	err := dbmap.TruncateTables()
	checkErr(err, "Trucate failed")

	// insert user
	u1 := models.User{0, "user1"}
	u2 := models.User{1, "user2"}

	err = dbmap.Insert(&u1, &u2)
	checkErr(err, "Insert failed")

	// fetch all rows
	var users []models.User
	_, err = dbmap.Select(&users, "select * from users order by user_id")
	checkErr(err, "Select failed")
	log.Printf("All rows:")
	for i, u := range users {
		log.Printf(" %d: %v\n", i, u)
	}

	return c.Render()
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
