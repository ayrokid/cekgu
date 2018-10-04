package controllers

import (
	"net/http"
	"strconv"

	"github.com/ayrokid/cekgu/models"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetQuestion(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	var question models.Question
	var response gin.H
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&question).Error
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "data found",
			"data":    question,
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (idb *InDB) AllQuestion(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	var (
		questions []models.Question
		response  gin.H
	)

	idb.DB.Find(&questions)
	if len(questions) <= 0 {
		response = gin.H{
			"message": "data not found",
			"data":    nil,
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "data found",
			"data":    questions,
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (idb *InDB) CreateQuestion(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	var (
		question models.Question
		response gin.H
	)

	err := c.ShouldBindJSON(&question)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		err = idb.DB.Create(&question).Error
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

func (idb *InDB) UpdateQuestion(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		question    models.Question
		newQuestion models.Question
		response    gin.H
	)

	err := idb.DB.Where("id = ? ", id).First(&question).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	}

	err = c.ShouldBindJSON(&newQuestion)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	}

	err = idb.DB.Model(&question).Update(newQuestion).Error
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "update successfully",
			"status":  true,
			"id":      id,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (idb *InDB) DeleteQuestion(c *gin.Context) {
	role, _ := c.Get("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "access not allow",
			"status":  false,
		})
		return
	}
	var (
		question models.Question
		response gin.H
	)
	id, _ := strconv.Atoi(c.Param("id"))
	err := idb.DB.Where("id = ? ", id).First(&question).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	} else {
		err = idb.DB.Delete(&question).Error
		if err != nil {
			response = gin.H{
				"message": "delete failed",
				"status":  false,
			}
		} else {
			response = gin.H{
				"message": "data deleted",
				"status":  true,
			}
		}
	}

	c.JSON(http.StatusOK, response)
}
