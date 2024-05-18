package database

import (
	"go-nicommerce/config"
	"go-nicommerce/models"
)

func InsertShipmentOptionByStoreID(shipmentOption *models.ShipmentOption) (interface{}, error) {

	if err := config.DB.Create(&shipmentOption).Error; err != nil {
		return nil, err
	}

	return shipmentOption, nil
}

func GetShipmentOptionsByStoreID(storeID int) (interface{}, error) {
	var shipmentOptions []models.ShipmentOption

	if err := config.DB.Where("store_id = ?", storeID).Find(&shipmentOptions).Error; err != nil {
		return nil, err
	}

	return shipmentOptions, nil
}

func GetShipmentOptionByID(id int) (*models.ShipmentOption, error) {
	var shipmentOption models.ShipmentOption

	if err := config.DB.Where("id = ?", id).Find(&shipmentOption).Error; err != nil {
		return nil, err
	}

	return &shipmentOption, nil
}
