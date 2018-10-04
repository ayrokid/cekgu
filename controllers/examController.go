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
	limit, _ := strconv.Atoi(c.Param("limit"))
	if limit <= 0 {
		limit = 5
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
	if exam.FinishDate.Before(tn) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Try out selesai",
			"status":  true,
		})
		return
	}

	/**
	 * Tampilan soal sebanyak 5
	 */
	dataQuest := []models.Question{}
	error = idb.DB.Limit(limit).Where("test_id = ?", testId).Find(&dataQuest).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
		})
		return
	}

	choice := models.Choice{}
	qc := models.QuestionChoice{}
	dataList := []models.QuestionChoice{}
	for _, element := range dataQuest {
		_ = idb.DB.Where("question_id = ? ", element.ID).Find(&choice).Error
		qc.QuestionId = element.ID
		qc.Content = element.Content
		qc.DataChoice = choice
	}
}
