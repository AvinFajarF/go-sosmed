package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AvinFajarF/internal/config"
	"github.com/AvinFajarF/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *gin.Context)  {
	
	// mengambil cookie
	tokenString := c.GetHeader("Authorization")


	// cek token

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
		// check exp token
		if (float64(time.Now().Unix()) > claims["exp"].(float64)) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user entity.Users



		log.Println(claims["sub"])
		log.Println(claims["exp"])

		config.DB.Where("id = ?", claims["sub"]).First()

		
		initializers.DB.First(&user, claims["sub"])

		// if user.ID == 0 {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }

		c.Set("user", user)
		c.Set("id", user.ID)

		c.Next()

	} else {
		log.Println(claims["sub"])
		log.Println(claims["exp"])
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}