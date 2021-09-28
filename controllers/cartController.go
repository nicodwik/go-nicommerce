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
	shippingPrice, _ := strconv.Atoi(c.FormValue("shipping_price"))

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
	var weight int
	if product.Weight < 1000 {
		weight = 1
	} else {
		weight = product.Weight / 1000
	}

	cart.ShippingPrice += (weight * qty) * shippingPrice
	cart.Discount += product.PriceCut * qty
	cart.TotalPrice += (product.BasePrice * qty) - cart.Discount + shippingPrice

	// insert cart detail
	_, err1 := database.InsertProductToCart(&cartDetail)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err1.Error(),
		})
	}

	_, err2 := database.UpdateCartInfo(cart)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	cartDetails, err := database.GetCartDetailsByCartID(cartID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	cart.CartDetails = cartDetails

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    cart,
	})
}

func DeleteProductFromCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))
	shippingPrice, _ := strconv.Atoi(c.FormValue("shipping_price"))
	cartDetailID, _ := strconv.Atoi(c.QueryParam("cart_detail_id"))

	// get cart detail
	cartDetail, err := database.GetCartDetailByID(cartDetailID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get product
	product, err := database.GetProductByID(int(cartDetail.ProductID))
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

	// delete cart detail
	_, err1 := database.DeleteProductFromCart(cartDetailID)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err1.Error(),
		})
	}

	// update cart
	var weight int
	if product.Weight < 1000 {
		weight = 1
	} else {
		weight = product.Weight / 1000
	}

	cart.ShippingPrice -= (weight * cartDetail.Qty) * shippingPrice
	cart.Discount -= product.PriceCut * cartDetail.Qty
	cart.TotalPrice -= (product.BasePrice * cartDetail.Qty) - cart.Discount + shippingPrice
	_, err2 := database.UpdateCartInfo(cart)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	// get cart details
	cartDetails, err := database.GetCartDetailsByCartID(cartID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	cart.CartDetails = cartDetails

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    cart,
	})
}
