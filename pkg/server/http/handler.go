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

	// mendapatkan request user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// generete hash
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPasswordString := string(hash)

	// proses register user
	result, err := h.userservice.Register(user.Username, hashedPasswordString, user.Email, user.Image, user.Bio)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}

	// return json response
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"result": result,
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {

	var user entity.Users

	// mendapatkan request user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "password atau email yang anda berikan salah", "status": "error"})
		return
	}

	// proses login user
	users, err := h.userservice.Login(user.Email, user.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "password atau email yang anda berikan salah", "status": "error"})
		return
	}

	// mendapatakan value SECRET dari file .env variable
	key := []byte(os.Getenv("SECRET"))

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"sub": users.ID,
	})

	tokenString, err := token.SignedString(key)

	c.SetSameSite(http.SameSiteLaxMode)

	// set token in headers Authorization
	c.Header("Authorization", tokenString)

	if err != nil {
		c.JSON(400, gin.H{"error": "password atau email yang anda berikan salah", "status": "error"})
		return
	}

	// return json response
	c.JSON(http.StatusOK, gin.H{
		"status": "oke",
	})

}

func (p *PostHandler) CreatePost(c *gin.Context) {
	// mendapatkan id dari user yang sudah login
	userId, _ := c.Get("id")

	userIdString, _ := userId.(string)

	var posts *entity.Posts

	// mendapatkan request user
	if err := c.ShouldBindJSON(&posts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// proses create post user
	result, err := p.postservice.CreatePosts(posts.Title, posts.Description, userIdString)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}

	// return json response
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   result,
	})

}

func (p *PostHandler) GetPosts(c *gin.Context) {
	// mendapatkan id dari user yang sudah login
	userId, _ := c.Get("id")

	userIdString, _ := userId.(string)

	// proses get post user
	result, err := p.postservice.GetPosts(userIdString, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	}
	// return json response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})

}

func (p *PostHandler) DeletePost(c *gin.Context) {

	// mendapatkan id dari user yang sudah login
	userId, _ := c.Get("id")
	userIdString, _ := userId.(string)

	// mendapatkan value dari parameter
	postId := c.Param("id")

	// mengecek apakah id dari user terdaftar atau tidak
	if err := p.postservice.FindUserById(userIdString, c); err != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	// proses delete user
	if error := p.postservice.DeletePost(postId); error != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "Error ketika menghapus post",
		})
		return
	}

	// return json response
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success ketika menghapus post",
	})

}

func (p *PostHandler) UpdatePost(c *gin.Context) {
	// mendapatkan id dari user yang sudah login
	userId, _ := c.Get("id")
	userIdString, _ := userId.(string)

	// mendapatkan value dari parameter
	postId := c.Param("id")

	// mendapatkan request
	var posts *entity.Posts

	if err := c.ShouldBindJSON(&posts); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// mengecek apakah id dari user terdaftar atau tidak
	if err := p.postservice.FindUserById(userIdString, c); err != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	// proses update user
	if error := p.postservice.UpdatePost(postId, posts.Title, posts.Description); error != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "Error ketika update post",
		})
		return
	}

	// return json response
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success ketika update post",
	})

}
