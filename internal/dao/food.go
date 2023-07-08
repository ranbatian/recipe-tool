package dao

import (
	"simon.com/recipe/global"
	"simon.com/recipe/internal/model"
)

type foodDao struct{}

func (*foodDao) Create(food *model.Food) error {
	err := global.G_DB.Create(food).Error
	return err
}

func (*foodDao) CreateUnit(food_unit *model.FoodUnit) error {
	err := global.G_DB.Create(food_unit).Error
	return err
}
