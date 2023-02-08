package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `json:"title" validate:"alpha"`
	Body  string `validate:"required"`
}
