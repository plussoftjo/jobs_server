// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// YoutubePosts ..
type YoutubePosts struct {
	gorm.Model
	Title string `json:"title"`
	Image string `json:"image"`
	YTID  string `json:"ytid"`
}
