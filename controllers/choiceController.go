package controllers

import (
	"net/http"
	"strconv"

	"github.com/ayrokid/cekgu/models"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetChoice(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
			"role":    role,
		})
		return
	}
	var choice []models.Choice
	var response gin.H
	q, _ := strconv.Atoi(c.Param("q"))
	err := idb.DB.Where("question_id = ?", q).Find(&choice).Error
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
		question models.Question
		response gin.H
	)
	q, _ := strconv.Atoi(c.Param("q"))
	err := idb.DB.Where("id = ? ", q).First(&question).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "data question not found",
			"status":  false,
		})
		return
	}

	err = c.ShouldBindJSON(&choice)
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
	q, _ := strconv.Atoi(c.Param("q"))
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		choice    models.Choice
		newChoice models.Choice
		response  gin.H
	)

	err := idb.DB.Where("id = ? AND question_id = ? ", id, q).First(&choice).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	}

	err = c.ShouldBindJSON(&newChoice)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	}

	err = idb.DB.Model(&choice).Update(&newChoice).Error
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

func (idb *InDB) DeleteChoice(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	q, _ := strconv.Atoi(c.Param("q"))
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		choice   models.Choice
		response gin.H
	)

	err := idb.DB.Where("id = ? AND question_id = ? ", id, q).First(&choice).Error
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

	err = idb.DB.Where("id = ? AND question_id = ? ", id, q).Delete(&choice).Error

	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "delete successfully",
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}
