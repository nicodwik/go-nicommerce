package database

import (
	"go-nicommerce/config"
	"go-nicommerce/models"
)

func InsertStore(store *models.Store) (interface{}, error) {
	if err := config.DB.Create(&store).Error; err != nil {
		return nil, err
	}

	return store, nil
}

func GetStoreByID(storeID int) (*models.Store, error) {
	var store models.Store

	if err := config.DB.First(&store).Error; err != nil {
		return nil, err
	}

	return &store, nil
}
