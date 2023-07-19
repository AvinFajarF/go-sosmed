package config

import (
	"log"
	"github.com/AvinFajarF/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func ConnectToDB() {

	dsn := "host=localhost user=postgres password=root dbname=go-sosmed port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	DB.AutoMigrate(model.Users{}, model.Posts{}, model.Like{}, model.Comments{})


}