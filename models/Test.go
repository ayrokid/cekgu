package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Test struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Duration    int        `json:"duration"`
	OpenDate    *time.Time `json:"open_date,omitempty"`
	CloseDate   *time.Time `json:"close_date,omitempty"`
	ByUserId    int        `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
