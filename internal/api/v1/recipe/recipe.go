package recipe

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
	var recipeDetail request.RecipeBody
	err := c.ShouldBind(&recipeDetail)
	if err != nil {
		global.G_LOG.Error("创建食谱参数错误")
		respond.FailureWithMsg("参数错误", c)
		return
	}
	service.RecipeService.CreateRecipe(recipeDetail)
}
