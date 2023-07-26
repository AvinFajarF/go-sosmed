package server

import (
	zx "net/http"

	"github.com/AvinFajarF/internal/middleware"
	"github.com/AvinFajarF/pkg/server/http"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *http.UserHandler, postHandler *http.PostHandler) *gin.Engine {

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", userHandler.RegisterUser)
		v1.POST("/login", userHandler.LoginUser)
		// v1.POST("/tes/:id", middleware.AuthMiddleware, Coba)
		// router untuk create post
		v1.POST("/post", middleware.AuthMiddleware, postHandler.CreatePost)
		// router untuk mendapatkan semua posts bedasarkan user yang di buat user
		v1.GET("/posts", middleware.AuthMiddleware, postHandler.GetPosts)
		// router untuk update post
		v1.PUT("/post/update/:id", middleware.AuthMiddleware, postHandler.UpdatePost)
		// router untuk delete post
		v1.DELETE("/post/delete/:id", middleware.AuthMiddleware, postHandler.DeletePost)
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
