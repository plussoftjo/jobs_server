// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// PostsRequest ..
type PostsRequest struct {
	gorm.Model
	Title       string `json:"title"`
	Body        string `json:"body"`
	Phone       string `json:"phone"`
	Place       string `json:"place"`
	Nationality string `json:"nationality"`
	Time        string `json:"time"`
	Uri         string `json:"uri"`
	UserID      uint   `json:"userID"`
	User        User   `json:"user"`
}
