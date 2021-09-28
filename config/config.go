package config

import (
	"fmt"
	"mini-project-acp12/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Host":     "127.0.0.1",
		"DB_Port":     "3306",
		"DB_Name":     "mini_project_acp12",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])
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
		// &models.User{},
		// &models.AddressOption{},
		// &models.Cart{},
		// &models.CartDetail{},
		// &models.Store{},
		// &models.ShipmentOption{},
		// &models.Category{},
		// &models.Product{},
		&models.ProductGallery{},
		// &models.Transaction{},
		// &models.TransactionProduct{},
	)
}
