package main

import (
	"log"

	"github.com/AvinFajarF/internal/model"
	"github.com/AvinFajarF/internal/repository"
	"github.com/AvinFajarF/internal/service"
	"github.com/AvinFajarF/pkg/server"
	"github.com/AvinFajarF/pkg/server/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=go-sosmed port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	db.AutoMigrate(model.Users{}, model.Posts{}, model.Like{}, model.Comments{})

	userRepository := repository.NewPostgreUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserService(&userService)

	router := server.NewRouter(userHandler)

	router.Run()
}
