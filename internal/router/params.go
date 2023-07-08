package router

import (
	"github.com/gin-gonic/gin"
	"simon.com/recipe/internal/api/v1/food"
	"simon.com/recipe/internal/api/v1/params"
	"simon.com/recipe/internal/api/v1/recipe"
)

func InitParamsRouter(group gin.RouterGroup) {
	paramsGroup := group.Group("/params")
	paramsGroup.GET("", params.List)
	paramsGroup.POST("", params.Create)
}

func InitRecipeRouter(group gin.RouterGroup) {
	recipeGroup := group.Group("/recipe")
	recipeGroup.GET("", recipe.List)
	recipeGroup.POST("", recipe.Create)
}

func InitFoodRouter(group gin.RouterGroup) {
	recipeGroup := group.Group("/food")
	recipeGroup.GET("/food_detail", food.List)
	recipeGroup.POST("/food_detail", food.Create)
	recipeGroup.GET("/food_unit", food.ListUnit)
	recipeGroup.POST("/food_unit", food.CreateUnit)
}
