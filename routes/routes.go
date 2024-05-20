package routes

import (
	"go-nicommerce/controllers"
	"go-nicommerce/env"

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

	// Category
	r.GET("/category/:store_id", controllers.GetCategoriesByStoreIDController)

	// Product
	r.GET("/product/:store_id", controllers.GetProductsByStoreIDController)

	// Transaction Callback (PG)
	r.POST("/transaction/callback", controllers.PaymentCallback)

	// route with middleware
	m := r.Group("/member")
	jwtSecret := env.Find("SECRETE_JWT", "legal")
	m.Use(middleware.JWT([]byte(jwtSecret)))

	// User
	m.GET("/users", controllers.GetUserController)
	m.GET("/users/:id", controllers.GetUserByIDController)
	m.PUT("/users", controllers.UpdateUserController)

	// Activate Store
	m.POST("/activate-store", controllers.ActivateStoreController)

	// Category
	m.POST("/category/:store_id", controllers.InsertCategoryController)

	// Product
	m.POST("/product/category/:category_id", controllers.InsertProductByCategoryIDController)
	m.PUT("/product/:product_id", controllers.UpdateProductByProductID)
	m.DELETE("/product/:product_id", controllers.DeleteProductByProductID)

	// Address Option
	m.GET("/address-option", controllers.GetAddressOptionsByUserIDController)
	m.POST("/address-option", controllers.InsertAddressOptionByUserIDController)

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
