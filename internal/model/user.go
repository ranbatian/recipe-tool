package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Age    int    `json:"age"`
}
