package config

import (
	"os"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	// db.AutoMigrate(models.Test{})
	// db.AutoMigrate(models.Question{})
	// db.AutoMigrate(models.Choice{})
	// db.AutoMigrate(models.User{})
	// db.AutoMigrate(models.Exam{})
	// db.AutoMigrate(models.Answer{})
	return db
}
