package models

import (
	"github.com/jinzhu/gorm"
)

type Dummy struct {
	gorm.Model
	SomeField string `json:"some_field"`
}
