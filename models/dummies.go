package models

import (
	"time"
)

type Dummy struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	SomeField string     `json:"some_field" binding:"required"`
}
