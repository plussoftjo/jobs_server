// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// User ..
type User struct {
	gorm.Model
	Name        string `json:"name"`
	Phone       string `json:"phone" gorm:"unique"`
	Password    string `json:"password"`
	RolesID     uint   `json:"roles_id"`
	UserType    uint   `json:"user_type"` // 01 -> User , 02 -> Controller
	PhoneCode   string `json:"phoneCode"`
	Nationality string `json:"nationality"`
	Avatar      string `json:"avatar"`
	Roles       Roles  `json:"roles" gorm:"foreignKey:RolesID;references:ID"`
	Block       int    `json:"block" gorm:"default:0"` // 0 => Active, 01 => Block
}

// Login ...
type Login struct {
	Phone    string `json:"phone" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}
