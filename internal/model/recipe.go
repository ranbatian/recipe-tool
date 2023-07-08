package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Describe string `json:"describe" gorm:"type:longtext"`
	FoodCommon
}
