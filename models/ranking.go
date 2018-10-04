package models

import (
	"github.com/jinzhu/gorm"
)

type Ranking struct {
	gorm.Model
	Score int    `json:"score"`
	Name  string `json:"name"`
}
