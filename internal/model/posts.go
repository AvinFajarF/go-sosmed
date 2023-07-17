package model

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Id          string     `gorm:"primaryKey"`
	UserID      uint       `gorm:"foreignKey:UserID"`
	User        Users 	   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Title       string
	Description string
}
