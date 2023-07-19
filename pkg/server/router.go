package server

import (
	zx "net/http"

	"github.com/AvinFajarF/internal/middleware"
	"github.com/AvinFajarF/pkg/server/http"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *http.UserHandler) *gin.Engine {

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", userHandler.RegisterUser)
		v1.POST("/login", userHandler.LoginUser)
		v1.POST("/tes/:id", middleware.AuthMiddleware, Coba)
	}

	return router
}

func Coba(c *gin.Context) {
	user, _ := c.Get("id")

	// Mendapatkan ID dari user
	c.JSON(zx.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
