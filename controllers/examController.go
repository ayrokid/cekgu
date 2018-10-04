package controllers

import (
	"cekgu/models"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) StartHandler(c *gin.Context) {
	testId, _ := strconv.Atoi(c.Param("id"))
	userId, _ := c.Get("UserID")

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
	page, _ := strconv.Atoi(c.Query("page"))
	offset := page * limit
	if page <= 0 {
		offset = 0
	}

	userId, _ := c.Get("UserID")

	error := idb.DB.Where("user_id = ? AND test_id = ? ", userId, testId).Select("finish_date").First(&exam).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
			"exam":    false,
		})
		return
	}

	/**
	 * Cek batas waktu ujian try out
	 */
	tn := time.Now()
	if exam.FinishDate.Before(tn) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Waktu try out habis",
			"status":  true,
		})
		return
	}

	/**
	 * Tampilan soal sebanyak 5
	 */
	dataQuest := []models.Question{}
	error = idb.DB.Limit(limit).Offset(offset).Where("test_id = ?", testId).Select("id, content").Find(&dataQuest).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
			"limit":   false,
		})
		return
	}

	choice := []models.Choice{}
	// qc := models.QuestionChoice{}
	dataList := []models.QuestionChoice{}
	for _, element := range dataQuest {
		_ = idb.DB.Where("question_id = ? ", element.ID).Select("id, choice").Find(&choice).Error
		qc := models.QuestionChoice{
			QuestionId: element.ID,
			Content:    element.Content,
			DataChoice: choice,
		}

		dataList = append(dataList, qc)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully",
		"status":  true,
		"data":    dataList,
	})
}

func (idb *InDB) AnswerHandler(c *gin.Context) {
	var exam models.Exam
	var answer models.Answer

	testId := c.Param("id")
	userId, _ := c.Get("UserID")

	error := idb.DB.Where("user_id = ? AND test_id = ? ", userId, testId).Select("finish_date").First(&exam).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
			"exam":    false,
		})
		return
	}

	/**
	 * Cek batas waktu ujian try out
	 */
	tn := time.Now()
	if exam.FinishDate.Before(tn) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Waktu try out habis",
			"status":  true,
		})
		return
	}

	err := c.ShouldBindJSON(&answer)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	err = idb.DB.Create(&answer).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "insert answer success",
		"status":  true,
	})
}

func (idb *InDB) FinishHandler(c *gin.Context) {
	var exam models.Exam

	testId := c.Param("id")
	userId, _ := c.Get("UserID")

	error := idb.DB.Where("user_id = ? AND test_id = ? ", userId, testId).Select("finish_date").First(&exam).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
			"exam":    false,
		})
		return
	}

	/**
	 * Cek batas waktu ujian try out
	 */
	tn := time.Now()
	if exam.FinishDate.Before(tn) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Waktu try out habis",
			"status":  true,
		})
		return
	}

	/**
	 * Hitung score
	 */
	var (
		correct     = 0
		wrong       = 0
		notAnswered = 0
	)
	question := models.Question{}
	listAnswer := []models.Answer{}
	error = idb.DB.Where("user_id = ? AND test_id = ? ", userId, testId).Select("id, response").Find(&listAnswer).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
			"answer":  false,
		})
		return
	}

	for _, answer := range listAnswer {
		_ = idb.DB.Where("question_id = ? ", answer.QuestionId).Select("answer").Find(&question).Error
		if answer.Response == question.Answer {
			correct += 1
		} else if answer.Response != question.Answer {
			wrong += 1
		}
	}

	var count int
	listQuestion := []models.Question{}
	_ = idb.DB.Where("test_id = ? ", testId).Find(&listQuestion).Count(&count)
	notAnswered = count - (correct + wrong)
	score := (correct * 4) + (wrong * -2)
	percenCorrect := math.Round(float64(correct)/float64(count)) * 100
	percenWrong := math.Round(float64(wrong)/float64(count)) * 100
	duration := time.Since(exam.StartDate)

	newExam := models.Exam{
		RightAnswer:       correct,
		WrongAnswer:       wrong,
		NotAnswered:       notAnswered,
		Score:             score,
		PercentageCorrect: percenCorrect,
		PercentageWrong:   percenWrong,
		FinishDate:        time.Now(),
		Duration:          duration.Minutes(),
		Status:            "completed",
	}

	error = idb.DB.Model(&exam).Update(newExam).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
		})
		return
	}

	/**
	 * Get rangking
	 */
	listExam := []models.Exam{}
	error = idb.DB.Where("test_id = ? ", testId).Order("Score desc").Select("user_id, score").Find(&listExam).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
		})
		return
	}

	position := 0
	for index, ele := range listExam {
		if userId == ele.UserId {
			position = int(index) + 1
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "succesfully",
		"status":   true,
		"position": position,
	})
}

func (idb *InDB) RankingHandler(c *gin.Context) {
	testId := c.Param("id")

	rangking := []models.Ranking{}
	error := idb.DB.Table("exams").Where("test_id = ? ", testId).Joins("left join users on users.id=exams.user_id").Order("Score desc").Select("score, name").Find(&rangking).Error
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": error.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "succesfully",
		"status":   true,
		"position": rangking,
	})
}
