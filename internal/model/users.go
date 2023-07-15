package model

import "gorm.io/gorm"

type UsersModel struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	Image    string
}
