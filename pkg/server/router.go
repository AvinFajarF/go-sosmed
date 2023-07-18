package server

import (
	"github.com/AvinFajarF/pkg/server/http"
	"github.com/gin-gonic/gin"
)


func NewRouter(userHandler *http.UserHandler) *gin.Engine {

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", userHandler.RegisterUser)
		v1.POST("/login", userHandler.LoginUser)
	}

	return router
}