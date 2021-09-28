package database

import (
	"mini-project-acp12/config"
	"mini-project-acp12/models"
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
