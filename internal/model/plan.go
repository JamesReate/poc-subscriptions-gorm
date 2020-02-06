package model

import "github.com/jinzhu/gorm"

type Plan struct {
	gorm.Model
	Name string
}
