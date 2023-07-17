package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	Id     string     `gorm:"primary key"`
	UserID uint       `gorm:"foreignKey:UserID"`
	User   Users 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	PostID uint       `gorm:"foreignKey:PostID"`
	Post   Posts 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"post"`
	Count  int
}
