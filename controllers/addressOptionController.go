package controllers

import (
	"mini-project-acp12/lib/database"
	"mini-project-acp12/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func InsertAddressOptionByUserIDController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	provinceID, _ := strconv.Atoi(c.FormValue("province_id"))
	cityID, _ := strconv.Atoi(c.FormValue("city_id"))

	addressOption := models.AddressOption{
		UserID:     uint(userID),
		ProvinceID: provinceID,
		CityID:     cityID,
		Address:    c.FormValue("address"),
	}

	savedAddressOption, err := database.InsertAddressOptionByUserID(&addressOption)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedAddressOption,
	})
}
