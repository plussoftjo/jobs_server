// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Tags ..
type Tags struct {
	gorm.Model
	Title string `json:"title"`
}
