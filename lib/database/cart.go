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

func DeleteProductFromCart(cartDetailID int) (interface{}, error) {
	var cartDetail models.CartDetail

	if err := config.DB.Where("id = ?", cartDetailID).Delete(&cartDetail).Error; err != nil {
		return nil, err
	}

	return cartDetail, nil
}

func GetCartDetailByID(cartDetailID int) (*models.CartDetail, error) {
	var cartDetail models.CartDetail

	if err := config.DB.Where("id = ?", cartDetailID).First(&cartDetail).Error; err != nil {
		return nil, err
	}

	return &cartDetail, nil
}

func GetCartByID(id int) (*models.Cart, error) {
	var cart models.Cart

	if err := config.DB.Where("id = ?", id).Preload("CartDetails").First(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func GetCartDetailsByCartID(id int) ([]models.CartDetail, error) {
	var cartDetails []models.CartDetail

	if err := config.DB.Where("cart_id = ?", id).Find(&cartDetails).Error; err != nil {
		return nil, err
	}

	return cartDetails, nil
}

func UpdateCartInfo(cart *models.Cart) (interface{}, error) {

	if err := config.DB.Save(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func EmptyCartByCartID(cartID int) error {

	if err := config.DB.Where("cart_id = ?", cartID).First(&cartID).Error; err != nil {
		return err
	}

	return nil
}
