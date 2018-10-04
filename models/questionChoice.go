package models

import (
	"github.com/jinzhu/gorm"
)

type QuestionChoice struct {
	gorm.Model
	QuestionId uint   `json:"id"`
	Content    string `json:"content"`
	DataChoice []Choice
}
