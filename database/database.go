package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadCredentials() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return map[string]string{
		"username":  os.Getenv("DB_USER"),
		"password":  os.Getenv("DB_PASSWORD"),
		"localhost": os.Getenv("DB_HOST"),
		"db_name":   os.Getenv("DB_NAME"),
		"port":      os.Getenv("DB_PORT"),
	}
}

func Init() *gorm.DB {
	credentials := loadCredentials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", credentials["username"], credentials["password"], credentials["localhost"], credentials["port"], credentials["db_name"])
	dbInstance, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		panic("Failed to connect to database")
	}

	DB = dbInstance
	fmt.Println("Connected to Database")
	return dbInstance
}
