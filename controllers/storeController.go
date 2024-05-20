package controllers

import (
	"go-nicommerce/lib/database"
	"go-nicommerce/middlewares"
	"go-nicommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ActivateStoreController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	province_id, _ := strconv.Atoi(c.FormValue("province_id"))
	city_id, _ := strconv.Atoi(c.FormValue("city_id"))

	store := models.Store{
		UserID:      uint(userId),
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Avatar:      c.FormValue("avatar"),
		ProvinceID:  province_id,
		CityID:      city_id,
		Address:     c.FormValue("address"),
	}

	savedStore, e := database.InsertStore(&store)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedStore,
	})
}
