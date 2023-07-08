package model

import (
	"gorm.io/gorm"
)

type FoodUnit struct {
	gorm.Model
	Desc string `json:"describe"`
	Name string `json:"name"`
}
