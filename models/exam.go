package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Exam struct {
	gorm.Model
	UserId     int64     `json:"user_id"`
	TestId     int       `json:"test_id"`
	StartDate  time.Time `json:"start_date"`
	FinishDate time.Time `json:"finish_date"`
	Duration   int       `json:"duration"`
	Score      int       `json:"score"`
	Status     string    `json:"status"`
}
