package models

import (
	"github.com/jinzhu/gorm"
)

// Key : store keys
type Key struct {
	gorm.Model

	Key    string  `gorm:"size:255;unique" json:"key"`
	Values []Value `json:"values"`
}

// TableName overrides the table name
func (Key) TableName() string {
	return "dragon_keys"
}
