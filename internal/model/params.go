package model

import "gorm.io/gorm"

type Params struct {
	gorm.Model
	FoodCommon
}
