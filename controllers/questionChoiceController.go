package controllers

import (
	"cekgu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetQuestionChoice(c *gin.Context) {
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

func (idb *InDB) AllQuestionChoice(c *gin.Context) {
	var (
		questionChoises []models.Choice
		response        gin.H
	)

	idb.DB.Find(&questionChoises)
	if len(questionChoises) <= 0 {
		response = gin.H{
			"message": "data not found",
			"data":    nil,
			"status":  false,
		}
	} else {
		response = gin.H{
			"message": "data found",
			"data":    questionChoises,
			"status":  true,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (idb *InDB) CreateQuestionChoice(c *gin.Context) {
	var (
		questionChoise models.QuestionChoice
		response       gin.H
	)

	err := c.ShouldBindJSON(&questionChoise)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		err = idb.DB.Create(&questionChoise).Error
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

func (idb *InDB) UpdateQuestionChoice(c *gin.Context) {
	id := c.Query("id")
	var (
		questionChoise    models.QuestionChoice
		newQuestionChoise models.QuestionChoice
		response          gin.H
	)

	err := idb.DB.First(&questionChoise, id).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	}

	err = c.ShouldBindJSON(&newQuestionChoise)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	}

	err = idb.DB.Model(&questionChoise).Update(newQuestionChoise).Error
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

func (idb *InDB) DeleteQuestionChoice(c *gin.Context) {
	var (
		questionChoise models.QuestionChoice
		response       gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&questionChoise, id).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	} else {
		err = idb.DB.Delete(&questionChoise).Error
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
