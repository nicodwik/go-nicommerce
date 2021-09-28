package controllers

import (
	"mini-project-acp12/lib/database"
	"mini-project-acp12/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func InsertProductGalleryByProductIDController(c echo.Context) error {
	productID, _ := strconv.Atoi(c.Param("product_id"))

	var isPrimary bool
	if c.FormValue("is_primary") == "true" {
		isPrimary = true
	} else {
		isPrimary = false
	}

	productGallery := models.ProductGallery{
		ProductID: uint(productID),
		Photo:     c.FormValue("photo"),
		IsPrimary: isPrimary,
	}

	savedProductGallery, err := database.InsertProductGalleryByProductID(&productGallery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedProductGallery,
	})
}

func GetProductGalleriesByProductIDController(c echo.Context) error {
	productID, _ := strconv.Atoi(c.Param("product_id"))

	productGalleries, err := database.GetProductGalleriesByProductID(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    productGalleries,
	})
}
