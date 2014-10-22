package controller

import (
	"github.com/mattn/kocho/db"
	"github.com/mattn/kocho/app/model"
	"github.com/naoina/kocha"
)

type Users struct {
	*kocha.DefaultController
}

func (us *Users) GET(c *kocha.Context) kocha.Result {
	var users []model.User
	err := db.Get("default").Select(&users)
	if err != nil {
		return kocha.RenderError(c, 500, err)
	}
	return kocha.RenderJSON(c, users)
}

func (us *Users) POST(c *kocha.Context) kocha.Result {
	name := c.Params.Get("name")
	if name == "" {
		return kocha.RenderError(c, 400, "Bad request")
	}
	user := &model.User{Name: name}
	_, err := db.Get("default").Insert(user)
	if err != nil {
		return kocha.RenderError(c, 500, err)
	}
	return kocha.RenderJSON(c, user)
}
