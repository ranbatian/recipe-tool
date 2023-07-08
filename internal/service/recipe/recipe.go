package recipe

import (
	"fmt"

	"simon.com/recipe/internal/model/common/request"
)

type RecipeService struct{}

func (*RecipeService) CreateRecipe(detail request.RecipeBody) {
	for _, food := range detail.FoodList {
		fmt.Printf("food: %v\n", food)
	}
}
