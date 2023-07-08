package core

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simon.com/recipe/global"
)

func InitDB() *gorm.DB {
	dsn := global.G_CONFIG.DB.CreateDNS()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	return db
}
