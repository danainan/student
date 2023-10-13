package database

import (
	"log"
	"os"
	"student/config"
	"student/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	dsn := config.Config("DATABASE_DNS")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Database has Connected")
	db.AutoMigrate(&models.Student{})
	DBConn = db
}
