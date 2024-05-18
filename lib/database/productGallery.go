package database

import (
	"go-nicommerce/config"
	"go-nicommerce/models"
)

func InsertProductGalleryByProductID(productGallery *models.ProductGallery) (interface{}, error) {
	if err := config.DB.Create(&productGallery).Error; err != nil {
		return nil, err
	}

	return productGallery, nil
}

func GetProductGalleriesByProductID(productID int) (interface{}, error) {
	var productGalleries []models.ProductGallery

	if err := config.DB.Where("product_id = ?", productID).Find(&productGalleries).Error; err != nil {
		return nil, err
	}

	return productGalleries, nil
}
