package controllers

import (
	"mini-project-acp12/lib/database"
	"mini-project-acp12/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func InsertProductToCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))
	productID, _ := strconv.Atoi(c.FormValue("product_id"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))

	cartDetail := models.CartDetail{
		CartID:    uint(cartID),
		ProductID: uint(productID),
		Qty:       qty,
	}

	// get product
	product, err := database.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get cart
	cart, err := database.GetCartByID(cartID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// update cart
	cart.Discount = product.PriceCut * qty
	cart.TotalPrice = (product.BasePrice * qty) - cart.Discount

	savedCartDetail, err := database.InsertProductToCart(&cartDetail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedCartDetail,
	})
}
