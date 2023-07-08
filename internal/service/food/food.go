package food

import (
	"simon.com/recipe/global"
	"simon.com/recipe/internal/dao"
	"simon.com/recipe/internal/model"
	"simon.com/recipe/internal/model/common/request"
)

type FoodService struct {
}

func (*FoodService) Create(req request.Food) error {
	food := model.Food{
		FoodCommon: model.FoodCommon{
			Fat:          req.Fat,
			Protein:      req.Protein,
			Carbohydrate: req.Carbohydrate,
		},
		Name: req.Name,
		Unit: req.FoodUint,
	}
	err := dao.FoodDao.Create(&food)
	if err != nil {
		global.G_LOG.Errorf("create food error. err %s", err.Error())
	}
	return err
}

func (*FoodService) CreateUnit(req request.FoodUnit) error {
	food := model.FoodUnit{
		Desc: req.Describe,
		Name: req.Name,
	}
	err := dao.FoodDao.CreateUnit(&food)
	if err != nil {
		global.G_LOG.Errorf("create food unit error. err %s", err.Error())
	}
	return err
}
