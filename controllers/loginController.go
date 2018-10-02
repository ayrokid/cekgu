package controllers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Credential struct {
	Username string `json:username`
	Password string `json:password`
}

func (idb *InDB) LoginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}

	if user.Username != "myname" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username",
		})
		c.Abort()
	} else if user.Password != "123123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong password",
		})
		c.Abort()
	} else {

		sign := jwt.New(jwt.GetSigningMethod("HS256"))
		token, err := sign.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
