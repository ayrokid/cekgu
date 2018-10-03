package main

import (
	"cekgu/config"
	"cekgu/controllers"
	"cekgu/middlewares"
	"os"

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

		v1.GET("/choice/:id", inDB.GetChoice)
	}

	tryout := router.Group("tryout")
	{
		tryout.POST("/start", auth, inDB.StartHandler)
	}

	router.Run(":" + port)
}
