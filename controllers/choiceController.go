package controllers

import (
	"cekgu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetChoice(c *gin.Context) {
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

func (idb *InDB) CreateChoice(c *gin.Context) {
	var (
		choise   models.Choice
		response gin.H
	)

	err := c.ShouldBindJSON(&choise)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		err = idb.DB.Create(&choise).Error
		if err != nil {
			response = gin.H{
				"message": "insert failed",
				"status":  false,
			}
		} else {
			response = gin.H{
				"message": "insert successfully",
				"status":  true,
			}
		}
	}

	c.JSON(http.StatusOK, response)

}

func (idb *InDB) UpdateChoice(c *gin.Context) {
	id := c.Query("id")
	var (
		choice    models.Choice
		newChoice models.Choice
		response  gin.H
	)

	err := idb.DB.First(&choice, id).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	}

	err = c.ShouldBindJSON(&choice)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	}

	err = idb.DB.Model(&choice).Update(newChoice).Error
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "update successfully",
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}
