package config

import (
	"fmt"
	"log"
	"mini-project-acp12/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// fmt.Println(connectionString)
	var e error
	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	DB = db
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.AddressOption{},
		&models.Cart{},
		&models.CartDetail{},
		&models.Store{},
		&models.ShipmentOption{},
		&models.Category{},
		&models.Product{},
		&models.ProductGallery{},
		&models.Transaction{},
		&models.TransactionProduct{},
	)
}
