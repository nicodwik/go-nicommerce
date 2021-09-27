package controllers

import (
	"fmt"
	"mini-project-acp12/lib/database"
	"mini-project-acp12/models"
	"net/http"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/labstack/echo"
)

func InsertProductByCategoryIDController(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("category_id"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))
	weight, _ := strconv.Atoi(c.FormValue("weight"))
	basePrice, _ := strconv.Atoi(c.FormValue("base_price"))
	priceCut, _ := strconv.Atoi(c.FormValue("price_cut"))

	category, err := database.GetCategoryByID(categoryId)
	fmt.Println(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	product := models.Product{
		CategoryID:  uint(categoryId),
		StoreID:     uint(category.StoreID),
		Name:        c.FormValue("name"),
		Slug:        slug.Make(c.FormValue("name")),
		Description: c.FormValue("description"),
		Stock:       stock,
		Weight:      weight,
		BasePrice:   basePrice,
		PriceCut:    priceCut,
	}

	savedProduct, err := database.InsertProductByCategoryID(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedProduct,
	})
}

func GetProductsByStoreIDController(c echo.Context) error {
	storeID, _ := strconv.Atoi(c.Param("store_id"))

	products, err := database.GetProductsByStoreID(storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    products,
	})

}
