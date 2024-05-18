package database

import (
	"go-nicommerce/config"
	"go-nicommerce/models"
)

func InsertProductByCategoryID(product *models.Product) (interface{}, error) {

	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func GetProductsByStoreID(storeID int) (interface{}, error) {
	var products []models.Product

	if err := config.DB.Where("store_id = ?", storeID).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByID(id int) (*models.Product, error) {
	var product models.Product

	if err := config.DB.Where("store_id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProductInfo(product *models.Product) (*models.Product, error) {

	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
