package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Mail     string `json:"mail" gorm:"unique"`
	Password string `json:"password"`
}
