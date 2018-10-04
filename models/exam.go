package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Exam struct {
	gorm.Model
	UserId            int64     `json:"user_id"`
	TestId            int       `json:"test_id"`
	StartDate         time.Time `json:"start_date"`
	FinishDate        time.Time `json:"finish_date"`
	Duration          float64   `json:"duration"`
	Score             int       `json:"score"`
	Status            string    `json:"status"`
	PercentageCorrect float64   `json:"percentage_correct"`
	PercentageWrong   float64   `json:"percentage_wrong"`
	NotAnswered       int       `json:"not_answered"`
	RightAnswer       int       `json:"right_answer"`
	WrongAnswer       int       `json:"wrong_answer"`
}
