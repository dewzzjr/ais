package model

import (
	"gorm.io/gorm"
)

type Article struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	gorm.Model
}

type Filter struct {
	Query  *string
	Author *string
}
