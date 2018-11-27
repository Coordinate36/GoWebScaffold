package models

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name  string `json:"name" binding:"required,max=128"`
	Age   uint8  `json:"age" binding:"required,gte=1,lte=130"`
	Email string `json:"email" binding:"required,email"`
}
