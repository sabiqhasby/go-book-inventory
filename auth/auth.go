package auth

import (
	"book-inventory/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})

}

func LoginPostHandler(c *gin.Context) {
	var credential models.Login

	err := c.Bind(&credential)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Username/password is invalid request",
		})
		return
	}

	if credential.Username != models.USER || credential.Password != models.PASSWORD {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Username/password is invalid",
		})
		return
	} else {

		// give token if valid
		claim := jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			Issuer:    "books inventory",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(models.SECRET))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"content": "token is invalid",
			})
			c.Abort()
		}

		//if valid then go to page books
		q := url.Values{}
		q.Set("auth", token)
		location := url.URL{Path: "/books", RawQuery: q.Encode()}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
	}

}
