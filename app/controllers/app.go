package controllers

import (
	"Go-Revel/app/models"
	"fmt"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// db格納と取得のサンプル
	// DbMap.Insert(&models.User{0, "user"})

	rows, _ := DbMap.Select(models.User{}, "select * from user")
	for _, row := range rows {
		user := row.(*models.User)
		fmt.Printf("%d, %s\n", user.Id, user.Name)
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
