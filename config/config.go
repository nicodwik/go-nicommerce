package config

import (
	"fmt"
	"go-nicommerce/env"
	"go-nicommerce/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	dbUsername := env.Find("DB_USERNAME", "root")
	dbPassword := env.Find("DB_PASSWORD", "root")
	dbHost := env.Find("DB_HOST", "host.docker.internal")
	dbPort := env.Find("DB_PORT", "3306")
	dbName := env.Find("DB_NAME", "go_nicommerce")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
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
