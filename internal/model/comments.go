package model

import "gorm.io/gorm"

type CommentsModel struct {
	gorm.Model
	Id      string     `gorm:"primary key"`
	UserID  uint       `gorm:"foreignKey:UserID"`
	User    UsersModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	PostID  uint       `gorm:"foreignKey:PostID"`
	Post    PostsModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"post"`
	Comment string
}
