package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	Image    string
	Bio      string
}
