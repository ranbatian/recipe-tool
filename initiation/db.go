package initization

import (
	"simon.com/recipe/global"
	"simon.com/recipe/internal/model"
)

func RegisterTable() error {
	err := global.G_DB.AutoMigrate(
		&model.Food{},
		&model.Recipe{},
		&model.User{},
		&model.Params{},
		&model.FoodUnit{},
	)
	return err
}
