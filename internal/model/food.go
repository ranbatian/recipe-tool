package model

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	FoodCommon
	Name string `json:"name"`
	Unit int    `json:"uint" gorm:"type:tinyint(4)"`
}
