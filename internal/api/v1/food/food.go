package food

import (
	"github.com/gin-gonic/gin"
	"simon.com/recipe/global"
	"simon.com/recipe/internal/model/common/request"
	"simon.com/recipe/internal/model/common/respond"
	"simon.com/recipe/internal/service"
)

func List(c *gin.Context) {

}

func Create(c *gin.Context) {
	var food request.Food
	err := c.ShouldBind(&food)
	if err != nil {
		global.G_LOG.Error("创建食物参数错误")
		respond.FailureWithMsg("创建食物参数错误", c)
		return
	}
	err = service.FoodService.Create(food)
	if err != nil {
		respond.FailureWithMsg("创建食物错误", c)
		return
	}
	respond.Success(c)
}

func CreateUnit(c *gin.Context) {
	var food_unit request.FoodUnit
	err := c.ShouldBind(&food_unit)
	if err != nil {
		global.G_LOG.Error("创建食物参数错误")
		respond.FailureWithMsg("创建食物参数错误", c)
		return
	}
	err = service.FoodService.CreateUnit(food_unit)
	if err != nil {
		respond.FailureWithMsg("创建食物单位错误", c)
		return
	}
	respond.Success(c)
}

func ListUnit(c *gin.Context) {

}
