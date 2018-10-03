package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
