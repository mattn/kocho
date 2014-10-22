package main

import (
	"github.com/mattn/kocho/app/model"
	"github.com/mattn/kocho/config"
	"github.com/mattn/kocho/db"

	"github.com/naoina/kocha"
)

func main() {
	db.Get("default").CreateTableIfNotExists(model.User{})

	if err := kocha.Run(config.AppConfig); err != nil {
		panic(err)
	}
}
