package params

import (
	"github.com/gin-gonic/gin"
	"simon.com/recipe/global"
	"simon.com/recipe/internal/model/common/request"
	"simon.com/recipe/internal/model/common/respond"
	"simon.com/recipe/internal/service"
)

func Create(c *gin.Context) {
	var params request.Params
	err := c.ShouldBind(&params)
	if err != nil {
		global.G_LOG.Error("参数输入错误")
		respond.Failure(c)
		return
	}
	err = service.Params.CreateParams(params)
	if err != nil {
		respond.Failure(c)
		return
	}
	respond.SuccessWithMsg("创建成功", c)
}

func List(c *gin.Context) {

}
