package database

import (
	"go-nicommerce/config"
	"go-nicommerce/models"
)

func InsertCategory(category *models.Category) (interface{}, error) {
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func GetCategoriesByStoreID(storeID int) (interface{}, error) {
	var categories []models.Category

	if err := config.DB.Where("store_id = ?", storeID).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryByID(id int) (*models.Category, error) {
	var category models.Category

	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
