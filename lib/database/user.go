package database

import (
	"mini-project-acp12/config"
	"mini-project-acp12/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Preload("Store").Preload("Cart").Find(&users).Error; err != nil { // Preload for retreive relationship data
		return nil, err
	}

	return users, nil
}

func InsertUser(user *models.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
