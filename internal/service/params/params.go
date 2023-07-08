package params

import (
	"simon.com/recipe/global"
	"simon.com/recipe/internal/dao"
	"simon.com/recipe/internal/model"
	"simon.com/recipe/internal/model/common/request"
)

type ParamsService struct{}

func (p *ParamsService) CreateParams(req request.Params) error {
	param := model.Params{
		FoodCommon: model.FoodCommon{
			Protein:      req.Protein,
			Fat:          req.Fat,
			Carbohydrate: req.Carbohydrate,
		},
	}
	err := dao.ParamsDao.Create(param)
	if err != nil {
		global.G_LOG.Errorf("创建参数错误, error %v", err)
	}
	return err
}
