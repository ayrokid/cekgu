package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	TestId   int    `json:"test_id"`
	Content  string `json:"content"`
	Answer   string `json:"answer"`
	ByUserId int    `json:"user_id"`
}
