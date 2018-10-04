package controllers

import (
	"cekgu/models"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (idb *InDB) LoginHandler(c *gin.Context) {
	var (
		user    models.User
		student models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	err = idb.DB.Where("email = ?", user.Email).First(&student).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			// "error": err.Error(),
			"message": "Password failed",
			"status":  false,
		})
		return
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["id"] = student.ID
	claims["role"] = student.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	sign.Claims = claims
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
