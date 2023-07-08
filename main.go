package main

import (
	"simon.com/recipe/global"
	initization "simon.com/recipe/initiation"
	"simon.com/recipe/internal/core"
)

func main() {
	global.G_LOG = initization.NewLog()
	global.G_VIPER = initization.InitConfig()
	global.G_DB = core.InitDB()
	if global.G_DB != nil {
		initization.RegisterTable()
		db, _ := global.G_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
