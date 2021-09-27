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

	// Category
	e.GET("/category/:store_id", controllers.GetCategoriesByStoreIDController)
	e.POST("/category/:store_id", controllers.InsertCategoryController)

	e.GET("/product/:store_id", controllers.GetProductsByStoreIDController)
	e.POST("/product/:category_id", controllers.InsertProductByCategoryIDController)

	e.POST("/address_option/:user_id", controllers.InsertAddressOptionByUserIDController)
	return e
}
