package controllers

import (
	"cekgu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) StartHandler(c *gin.Context) {
	// ext := middlewares.ExtractClaims
	c.JSON(http.StatusOK, gin.H{
		"message": "mulai gan",
		// "ext":     string(ext),
	})

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
