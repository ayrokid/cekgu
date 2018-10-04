package models

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	UserId     int    `json:"user_id"`
	TestId     int    `json:"test_id"`
	QuestionId int    `json:"question_id"`
	Response   string `json:"response"`
}
