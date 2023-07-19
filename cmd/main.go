package main

import (
	"log"

	"github.com/AvinFajarF/internal/config"
	"github.com/AvinFajarF/internal/repository"
	"github.com/AvinFajarF/internal/service"
	"github.com/AvinFajarF/pkg/server"
	"github.com/AvinFajarF/pkg/server/http"
	"github.com/joho/godotenv"
)




func init() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
}

func main() {
	
	userRepository := repository.NewPostgreUserRepository(config.DB)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserService(&userService)

	postRepository := repository.NewPostgrePostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := http.NewPostService(&postService)

	router := server.NewRouter(userHandler, postHandler)

	router.Run(":8081")
}
