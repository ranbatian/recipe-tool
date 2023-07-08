package service

import (
	"simon.com/recipe/internal/service/food"
	"simon.com/recipe/internal/service/params"
	"simon.com/recipe/internal/service/recipe"
)

var (
	Params        = params.ParamsService{}
	RecipeService = recipe.RecipeService{}
	FoodService   = food.FoodService{}
)
