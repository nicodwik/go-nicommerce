package database

import (
	"go-nicommerce/config"
	"go-nicommerce/middlewares"
	"go-nicommerce/models"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(email string, rawPassword string) (interface{}, error) {
	var err error
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	errr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rawPassword))
	if errr != nil {
		return nil, errr
	}

	user.Token, err = middlewares.CreateToken(int(user.ID.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Preload("Store").Preload("Cart").Find(&users).Error; err != nil { // Preload for retreive relationship data
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Preload("Cart").Preload("Store").Where("id = ?", id).First(&user).Error; err != nil { // Preload for retreive relationship data
		return nil, err
	}

	return &user, nil
}

func InsertUser(user *models.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
