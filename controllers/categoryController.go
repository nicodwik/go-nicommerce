package controllers

import (
	"go-nicommerce/lib/database"
	"go-nicommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func InsertCategoryController(c echo.Context) error {
	storeID, _ := strconv.Atoi(c.Param("store_id"))

	category := models.Category{
		StoreID: uint(storeID),
		Name:    c.FormValue("name"),
		Avatar:  c.FormValue("avatar"),
	}

	savedCategory, err := database.InsertCategory(&category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedCategory,
	})
}

func GetCategoriesByStoreIDController(c echo.Context) error {
	storeID, _ := strconv.Atoi(c.Param("store_id"))

	categories, err := database.GetCategoriesByStoreID(storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    categories,
	})
}
