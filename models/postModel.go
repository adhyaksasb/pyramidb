package model

import "gorm.io/gorm"

type Post struct {
	Title string
	Body string
	gorm.Model
}