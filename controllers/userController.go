package controllers

import (
	"mini-project-acp12/helpers"
	"mini-project-acp12/lib/database"
	"mini-project-acp12/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetUserController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func InsertUserController(c echo.Context) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), 14)
	if err != nil {
		return helpers.ServerErrorResponse(err.Error())
	}

	user := models.User{
		Firstname:   c.FormValue("firstname"),
		Lastname:    c.FormValue("lastname"),
		Email:       c.FormValue("email"),
		Phone:       c.FormValue("phone"),
		Avatar:      c.FormValue("avatar"),
		Password:    string(hashed),
		StoreStatus: false,
		Cart:        models.Cart{Timestamp: models.Timestamp{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
	}
	// user.Cart = models.Cart{UserID: user.ID.ID}

	savedCart, e := database.InsertUser(&user)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedCart,
	})
}
