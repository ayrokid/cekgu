package controllers

import (
	"guru/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetTest(c *gin.Context) {
	var test models.Test
	var response gin.H
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&test).Error
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "data found",
			"data":    test,
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (idb *InDB) allTest(c *gin.Context) {
	var (
		tests    []models.Test
		response gin.H
	)

	idb.DB.Find(&tests)
	if len(tests) <= 0 {
		response = gin.H{
			"message": "data not found",
			"data":    nil,
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "data found",
			"data":    tests,
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (idb *InDB) CreateTest(c *gin.Context) {
	var (
		test    models.Test
		message gin.H
	)

}
