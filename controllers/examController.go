package controllers

import (
	"cekgu/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) StartHandler(c *gin.Context) {
	testId, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": er,
			"status":  false,
			"id":      testId,
		})
		return
	}
	userId, err := c.Get("UserID")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err,
			"status":  false,
			"userId":  userId,
		})
		return
	}

	exam := models.Exam{
		UserId:     userId.(int64),
		TestId:     testId,
		StartDate:  time.Now(),
		FinishDate: time.Now().Local().Add(time.Minute * time.Duration(15)),
		Status:     "open",
	}
	idb.DB.Create(&exam)

	c.JSON(http.StatusOK, gin.H{
		"message": "mulai gan",
		"id":      userId,
		"exam":    exam,
	})

}

func (idb *InDB) ExamHandler(c *gin.Context) {
	var exam models.Exam
	testId := c.Param("id")

	userId, err := c.Get("UserID")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err,
			"status":  false,
			"userId":  userId,
		})
		return
	}

	error := idb.DB.Where("user_id = ? AND test_id = ? ", userId, testId).First(&exam).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
		})
		return
	}

	/**
	 * Cek batas waktu ujian try out
	 */
	tn := time.Now()
	statuUjian := "finised"
	if exam.FinishDate.After(tn) {
		statuUjian = "started"
	}

}
