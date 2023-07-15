package model

import "gorm.io/gorm"

type PostsModel struct {
	gorm.Model
	Id          string     `gorm:"primaryKey"`
	UserID      uint       `gorm:"foreignKey:UserID"`
	User        UsersModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Title       string
	Description string
}
