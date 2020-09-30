package models

import (
	"github.com/jinzhu/gorm"
)

// Value : tracks values of a key
type Value struct {
	gorm.Model

	Name  string `gorm:"size:255" json:"name"`
	KeyID uint   `json:"key_id"`
}

// TableName overrides the table name
func (Value) TableName() string {
	return "dragon_values"
}
