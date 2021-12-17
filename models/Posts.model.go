// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Posts ..
type Posts struct {
	gorm.Model
	Title       string `json:"title"`
	Body        string `json:"body"`
	Image       string `json:"image"`
	Phone       string `json:"phone"`
	Place       string `json:"place"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
	Time        string `json:"time"`
	Tags        string `json:"tags"`
	Uri         string `json:"uri"`
}
