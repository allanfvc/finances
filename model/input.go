package model

import (
	"github.com/jinzhu/gorm"
)

type Input struct {
	gorm.Model
	description string
	value       float32
	typeInput   int
	status      int
	user        User
	month       int
	year        int
}
