package db

import (
	"books_go/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	godotenv.Load()
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("No DB_URL found in environvent")
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("database connection successful")
	book := models.Book{}
	DB.Find(&book)
	fmt.Println(book)
	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		fmt.Println("Migration error")
		return err
	} else {
		fmt.Println("Migration successful")
		return nil
	}
}
