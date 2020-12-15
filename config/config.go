package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func connectDB(h *gorm.DB) {
	//connect database mysql by gorm
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FBangkok&parseTime=true"
	db, err := gorm.Open(dbConnection, dsn)
	if err != nil {
		log.Fatal("Database connect refuse")
	}
	db.Close()
}
