package models

import (
	"github.com/jinzhu/gorm"
)

type QuestionChoice struct {
	gorm.Model
	QuestionId int    `json:"question_id"`
	Choice     string `json:"choice"`
}
