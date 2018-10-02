package controllers

import (
	"cekgu/models"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (idb *InDB) StartHandler(c *gin.Context) {
	var (
		user     models.User
		response gin.H
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
		c.JSON(http.StatusUnauthorized, response)
	}

	email := user.Email
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err == nil {
		session := sessions.Default(c)
		session.Set("user", email)
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
		}
	}
}

func (idb *InDB) ExamHandler(c *gin.Context) {
	var choice models.Choice
	var response gin.H
	id := c.Param("id")
	err := idb.DB.Where("question_id = ?", id).Find(&choice).Error
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "data found",
			"data":    choice,
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}
