package request

type Params struct {
	Fat          float32 `json:"fat"`
	Carbohydrate float32 `json:"carbohydrate"`
	Protein      float32 `json:"protein"`
}

type RecipeBody struct {
	Name     string       `json:"name"`
	Describe string       `json:"describe"`
	FoodList []FoodDetail `json:"foodList"`
}

type FoodDetail struct {
	FoodId     int
	WeightUnit int
}

type Food struct {
	Name         string  `json:"name"`
	FoodUint     int     `json:"uint"`
	Fat          float32 `json:"fat"`
	Carbohydrate float32 `json:"carbohydrate"`
	Protein      float32 `json:"protein"`
}

type FoodUnit struct {
	Name     string `json:"name"`
	Describe string `json:"describe"`
}
