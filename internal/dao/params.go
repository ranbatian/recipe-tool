package dao

import (
	"simon.com/recipe/global"
	"simon.com/recipe/internal/model"
)

type paramsDao struct{}

func (*paramsDao) Create(param model.Params) error {
	err := global.G_DB.Create(&param).Error
	return err
}
