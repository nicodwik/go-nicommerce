package config

import (
	"fmt"
	"mini-project-acp12/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", constants.DB_USERNAME, constants.DB_PASSWORD, constants.DB_HOST, constants.DB_PORT, constants.DB_NAME)
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
	// 	&models.User{},
	// 	&models.AddressOption{},
	// 	&models.Cart{},
	// 	&models.CartDetail{},
	// 	&models.Store{},
	// 	&models.ShipmentOption{},
	// 	&models.Category{},
	// 	&models.Product{},
	// 	&models.ProductGallery{},
	// 	&models.Transaction{},
	// 	&models.TransactionProduct{},
	)
}
