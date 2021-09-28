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
	e.POST("/login", controllers.LoginController)

	// Store
	e.POST("/store/:user_id", controllers.ActivateStoreController)

	// Category
	e.GET("/category/:store_id", controllers.GetCategoriesByStoreIDController)
	e.POST("/category/:store_id", controllers.InsertCategoryController)

	// Product
	e.GET("/product/:store_id", controllers.GetProductsByStoreIDController)
	e.POST("/product/:category_id", controllers.InsertProductByCategoryIDController)

	// Address Option
	e.GET("/address-option/:user_id", controllers.GetAddressOptionsByUserIDController)
	e.POST("/address-option/:user_id", controllers.InsertAddressOptionByUserIDController)

	// Shipment Option
	e.GET("/shipment-option/:store_id", controllers.GetShipmentOptionsByStoreIDController)
	e.POST("/shipment-option/:store_id", controllers.InsertShipmentOptionByStoreIDController)

	// Product Gallery
	e.GET("/product-gallery/:product_id", controllers.GetProductGalleriesByProductIDController)
	e.POST("/product-gallery/:product_id", controllers.InsertProductGalleryByProductIDController)

	// Cart
	e.POST("/cart/:cart_id", controllers.InsertProductToCart)
	e.DELETE("/cart/:cart_id", controllers.DeleteProductFromCart)

	// Transaction
	e.POST("/transaction/:cart_id", controllers.InsertTransactionController)

	return e
}
