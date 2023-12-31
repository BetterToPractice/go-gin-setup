package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `gorm:"column:title;size:255;not null;" json:"title" validate:"required"`
	Body  string `gorm:"column:body;not null;" json:"body" validate:"required"`

	UserID uint `gorm:"column:user_id;not null;"`
	User   User `gorm:"foreignKey:UserID;references:ID;"`
}

type Posts []Post
