package database

import (
	"mini-project-acp12/config"
	"mini-project-acp12/models"
)

func InsertCartBelongsToUser(user *models.User) (interface{}, error) {
	var cart models.Cart

	cart.UserID = 23
	if err := config.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func InsertProductToCart(cartDetail *models.CartDetail) (interface{}, error) {

	if err := config.DB.Create(&cartDetail).Error; err != nil {
		return nil, err
	}

	return cartDetail, nil
}

func GetCartByID(id int) (*models.Cart, error) {
	var cart models.Cart

	if err := config.DB.Where("id = ?", id).First(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}
