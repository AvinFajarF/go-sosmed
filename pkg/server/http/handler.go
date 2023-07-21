package http

import (
	"net/http"
	"os"
	"time"

	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userservice service.UserService
}

type PostHandler struct {
	postservice service.PostService
}

func NewUserService(userservice service.UserService) *UserHandler {
	return &UserHandler{
		userservice: userservice,
	}
}

func NewPostService(postservice service.PostService) *PostHandler {
	return &PostHandler{
		postservice: postservice,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user entity.Users

	if err := c.ShouldBindJSON(&user); err != nil {
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
		"result": result,
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {

	var user entity.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "password atau email yang anda berikan salah", "status": "error"})
		return
	}

	users, err := h.userservice.Login(user.Email, user.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "password atau email yang anda berikan salah", "status": "error"})
		return
	}

	key := []byte(os.Getenv("SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"sub": users.ID,
	})

	tokenString, err := token.SignedString(key)

	c.SetSameSite(http.SameSiteLaxMode)
	c.Header("Authorization", tokenString)

	if err != nil {
		c.JSON(400, gin.H{"error": "password atau email yang anda berikan salah", "status": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "oke",
	})

}

func (p *PostHandler) CreatePost(c *gin.Context) {

	userId, _ := c.Get("id")

	userIdString, _ := userId.(string)


	var posts *entity.Posts

	if err := c.ShouldBindJSON(&posts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := p.postservice.CreatePosts(posts.Title, posts.Description, userIdString)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
        return
	}

	c.JSON(http.StatusCreated, gin.H{
        "status": "success",
        "data": result,
    })

}


func (p *PostHandler) GetPosts(c *gin.Context) {
	userId, _ := c.Get("id")

	userIdString, _ := userId.(string)
	
	result, err := p.postservice.GetPosts(userIdString, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
        "error": err.Error(),
        })
	}

	c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": result,
		})

}

func (p *PostHandler) DeletePost(c *gin.Context) {

}
