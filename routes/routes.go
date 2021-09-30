package routes

import (
	"mini-project-acp12/constants"
	"mini-project-acp12/controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	r := e.Group("/api/v1")

	// route without middleware
	r.POST("/login", controllers.LoginController)
	r.POST("/register", controllers.InsertUserController)

	// route with middleware
	m := r.Group("/member")
	m.Use(middleware.JWT([]byte(constants.SECRETE_JWT)))

	// User
	m.GET("/users", controllers.GetUserController)
	m.GET("/users/:id", controllers.GetUserByIDController)
	m.PUT("/users", controllers.UpdateUserController)

	// Store
	m.POST("/store/:user_id", controllers.ActivateStoreController)

	// Category
	m.GET("/category/:store_id", controllers.GetCategoriesByStoreIDController)
	m.POST("/category/:store_id", controllers.InsertCategoryController)

	// Product
	m.GET("/product/:store_id", controllers.GetProductsByStoreIDController)
	m.POST("/product/category/:category_id", controllers.InsertProductByCategoryIDController)

	// Address Option
	m.GET("/address-option/:user_id", controllers.GetAddressOptionsByUserIDController)
	m.POST("/address-option/:user_id", controllers.InsertAddressOptionByUserIDController)

	// Shipment Option
	m.GET("/shipment-option/:store_id", controllers.GetShipmentOptionsByStoreIDController)
	m.POST("/shipment-option/:store_id", controllers.InsertShipmentOptionByStoreIDController)

	// Product Gallery
	m.GET("/product-gallery/:product_id", controllers.GetProductGalleriesByProductIDController)
	m.POST("/product-gallery/:product_id", controllers.InsertProductGalleryByProductIDController)

	// Cart
	m.POST("/cart/:cart_id", controllers.InsertProductToCart)
	m.DELETE("/cart/:cart_id", controllers.DeleteProductFromCart)

	// Transaction
	m.POST("/transaction/:cart_id", controllers.InsertTransactionController)

	return e
}
