package controllers

import (
	"cekgu/models"
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

func (idb *InDB) AllTest(c *gin.Context) {
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
		test     models.Test
		response gin.H
	)

	err := c.ShouldBindJSON(&test)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	} else {
		err = idb.DB.Create(&test).Error
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

func (idb *InDB) UpdateTest(c *gin.Context) {
	id := c.Query("id")
	var (
		test     models.Test
		newTest  models.Test
		response gin.H
	)

	err := idb.DB.First(&test, id).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	}

	err = c.ShouldBindJSON(&newTest)
	if err != nil {
		response = gin.H{
			"message": err.Error(),
			"status":  false,
		}
	}

	err = idb.DB.Model(&test).Update(newTest).Error
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

func (idb *InDB) DeleteTest(c *gin.Context) {
	var (
		test     models.Test
		response gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&test, id).Error
	if err != nil {
		response = gin.H{
			"message": "data not found",
			"status":  false,
		}
	} else {
		err = idb.DB.Delete(&test).Error
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