package routes

import (
	"mini-project-acp12/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// User
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.InsertUserController)

	// Store
	e.POST("/store/:user_id", controllers.ActivateStoreController)
	return e
}
