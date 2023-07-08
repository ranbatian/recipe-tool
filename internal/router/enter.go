package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(eng *gin.Engine) {
	publicGroup := eng.Group("v1")
	{
		publicGroup.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
	privateGroup := eng.Group("v1")
	{
		InitParamsRouter(*privateGroup)
		InitFoodRouter(*privateGroup)
		InitRecipeRouter(*privateGroup)
	}
}
