package controllers

import (
	"go-nicommerce/helpers"
	"go-nicommerce/lib/database"
	"go-nicommerce/middlewares"
	"go-nicommerce/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	users, err := database.LoginUser(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

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

func UpdateUserController(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)

	hashed, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), 14)
	if err != nil {
		return helpers.ServerErrorResponse(err.Error())
	}

	user, _ := database.GetUserByID(userID)

	user.Firstname = c.FormValue("firstname")
	user.Lastname = c.FormValue("lastname")
	user.Email = c.FormValue("email")
	user.Phone = c.FormValue("phone")
	user.Avatar = c.FormValue("avatar")
	user.Password = string(hashed)

	savedUser, err := database.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  savedUser,
	})
}
