package database

import (
	"go-nicommerce/config"
	"go-nicommerce/models"
)

func InsertAddressOptionByUserID(addressOption *models.AddressOption) (interface{}, error) {

	if err := config.DB.Create(&addressOption).Error; err != nil {
		return nil, err
	}

	return addressOption, nil
}

func GetAddressOptionsByUserID(userID int) (interface{}, error) {
	var addressOptions []models.AddressOption

	if err := config.DB.Where("user_id = ?", userID).Find(&addressOptions).Error; err != nil {
		return nil, err
	}

	return addressOptions, nil
}

func GetAddressOptionByID(id int) (*models.AddressOption, error) {
	var addressOption models.AddressOption

	if err := config.DB.Where("id = ?", id).Find(&addressOption).Error; err != nil {
		return nil, err
	}

	return &addressOption, nil
}
