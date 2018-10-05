package main

import (
	"os"

	"cekgu/config"

	"github.com/ayrokid/cekgu/controllers"

	"github.com/ayrokid/cekgu/middlewares"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	auth := middlewares.Auth
	router := gin.Default()

	router.GET("/", inDB.HelloHandler)
	router.POST("/login", inDB.LoginHandler)
	v1 := router.Group("v1")
	{
		v1.GET("/test/:id", auth, inDB.GetTest)
		v1.GET("/tests", auth, inDB.AllTest)
		v1.POST("/test", auth, inDB.CreateTest)
		v1.PUT("/test/:id", auth, inDB.UpdateTest)
		v1.DELETE("/test/:id", auth, inDB.DeleteTest)

		v1.GET("/question/:id", auth, inDB.GetQuestion)
		v1.GET("/questions", auth, inDB.AllQuestion)
		v1.POST("/question", auth, inDB.CreateQuestion)
		v1.PUT("/question/:id", auth, inDB.UpdateQuestion)
		v1.DELETE("/question/:id", auth, inDB.DeleteQuestion)

		v1.GET("/choice/:q", auth, inDB.GetChoice)
		v1.POST("/choice/:q", auth, inDB.CreateChoice)
		v1.PUT("/choice/:q/:id", auth, inDB.UpdateChoice)
		v1.DELETE("/choice/:q/:id", auth, inDB.DeleteChoice)
	}

	tryout := router.Group("tryout")
	{
		tryout.GET("/start/:id", auth, inDB.StartHandler)
		tryout.GET("/exam/:id/:limit", auth, inDB.ExamHandler)
		tryout.POST("/answer/:id", auth, inDB.AnswerHandler)
		tryout.GET("/finish/:id", auth, inDB.FinishHandler)
		tryout.GET("/ranking/:id", auth, inDB.RankingHandler)
	}

	router.Run(":" + port)
}
