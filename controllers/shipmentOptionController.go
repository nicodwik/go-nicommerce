package controllers

import (
	"go-nicommerce/lib/database"
	"go-nicommerce/middlewares"
	"go-nicommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func InsertShipmentOptionController(c echo.Context) error {
	storeID := middlewares.ExtractTokenUserId(c)

	_, err := database.GetShipmentOptionsByStoreID(storeID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Store Not Found",
		})
	}

	shipmentOption := models.ShipmentOption{
		StoreID:  uint(storeID),
		Name:     c.FormValue("name"),
		IsActive: true,
		Avatar:   c.FormValue("avatar"),
	}

	savedShipmentOption, err := database.InsertShipmentOptionByStoreID(&shipmentOption)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedShipmentOption,
	})
}

func GetShipmentOptionsByStoreIDController(c echo.Context) error {
	storeID, _ := strconv.Atoi(c.Param("store_id"))

	shipmentOptions, err := database.GetShipmentOptionsByStoreID(storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    shipmentOptions,
	})
}
