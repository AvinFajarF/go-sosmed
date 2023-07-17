package http

import (
	"net/http"

	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userservice service.UserService
}

func NewUserService(userservice service.UserService)  *userHandler {
	return &userHandler{
		userservice: userservice,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var user entity.UserEntity

	if err := c.ShouldBindJSON(&user); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
		return
    }

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPasswordString := string(hash)

	result, err := h.userservice.Register(user.Username, hashedPasswordString, user.Email, user.Image, user.Bio)

	if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"user": result,
	})
}






