package database

import (
	"mini-project-acp12/config"
	"mini-project-acp12/middlewares"
	"mini-project-acp12/models"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(user *models.User, rawPassword string) (interface{}, error) {
	var err error

	if err := config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return nil, err
	}

	errr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rawPassword))
	if errr != nil {
		return nil, err
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
