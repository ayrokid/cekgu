package controllers

import (
	"cekgu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetChoice(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
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
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	var (
		choice   models.Choice
		response gin.H
	)

	err := c.ShouldBindJSON(&choice)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		err = idb.DB.Create(&choice).Error
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
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	id := c.Query("id")
	var (
		choice   models.Choice
		response gin.H
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

	_ = idb.DB.Where("question_id = ? ", id).Delete(&choice).Error

	err = idb.DB.Create(&choice).Error
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
