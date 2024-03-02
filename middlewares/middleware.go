package middlewares

import (
	"book-inventory/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthValid(c *gin.Context) {
	var tokenString string

	tokenString = c.Query("auth")
	if tokenString == "" {
		tokenString = c.PostForm("auth")
		if tokenString == "" {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "token nil"})
			c.Abort()
			return
		}
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, invalid := t.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("invalid token ", t.Header["alg"])
		}
		return []byte(models.SECRET), nil
	})

	if token == nil || err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "token is invalid"})
		c.Abort()
		return
	} else {
		fmt.Println("token verified")
		c.Next()
	}

}
