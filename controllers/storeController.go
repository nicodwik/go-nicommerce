package controllers

import (
	"mini-project-acp12/lib/database"
	"mini-project-acp12/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ActivateStoreController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
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
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedStore,
	})
}
