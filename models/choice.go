package models

import (
	"github.com/jinzhu/gorm"
)

type Choice struct {
	gorm.Model
	QuestionId int    `json:"question_id"`
	Choice     string `json:"choise"`
}
