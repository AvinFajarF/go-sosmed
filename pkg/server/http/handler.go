package http

import (
	"net/http"

	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userservice service.UserService
}

func NewUserService(userservice service.UserService)  *UserHandler {
	return &UserHandler{
		userservice: userservice,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user entity.Users

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
        c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"user": result,
	})
}






