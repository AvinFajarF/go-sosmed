package model

import "gorm.io/gorm"

type UsersModel struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	Image    string
	Bio      string
}
