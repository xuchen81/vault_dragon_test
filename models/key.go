package models

import (
	"github.com/jinzhu/gorm"
)

// Key : store keys
type Key struct {
	gorm.Model

	Name   string  `gorm:"size:255;unique" json:"name"`
	Values []Value `json:"values"`
}

// TableName overrides the table name
func (Key) TableName() string {
	return "dragon_keys"
}
