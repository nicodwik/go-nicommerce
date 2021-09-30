package controllers

import (
	"encoding/json"
	"io/ioutil"
	"mini-project-acp12/lib/database"
	"mini-project-acp12/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type shippingData struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Courier  string `json:"courier"`
}

type detailTransaction struct {
	transactionProducts []models.TransactionProduct
}

type province struct {
	Nama string `json:"nama"`
}

type city struct {
	IdProvinsi string `json:"id_provinsi"`
	Nama       string `json:"nama"`
}

func InsertTransactionController(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))
	storeID, _ := strconv.Atoi(c.FormValue("store_id"))
	shippingID, _ := strconv.Atoi(c.FormValue("shipping_id"))
	addressID, _ := strconv.Atoi(c.FormValue("address_id"))

	// get cart
	cart, err := database.GetCartByID(cartID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get address option
	addressOption, err := database.GetAddressOptionByID(addressID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get address option
	shipmentOption, err := database.GetShipmentOptionByID(shippingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get user
	user, err := database.GetUserByID(int(addressOption.UserID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get province data
	provinceResponse, _ := http.Get("https://dev.farizdotid.com/api/daerahindonesia/provinsi/" + strconv.Itoa(addressOption.ProvinceID))
	provinceResponseData, _ := ioutil.ReadAll(provinceResponse.Body)
	var province province
	json.Unmarshal(provinceResponseData, &province)

	// get city data
	cityResponse, _ := http.Get("https://dev.farizdotid.com/api/daerahindonesia/kota/" + strconv.Itoa(addressOption.CityID))
	cityResponseData, _ := ioutil.ReadAll(cityResponse.Body)
	var city city
	json.Unmarshal(cityResponseData, &city)

	shippingData := shippingData{
		Name:     user.Firstname + user.Lastname,
		Phone:    user.Phone,
		Province: province.Nama,
		City:     city.Nama,
		Address:  addressOption.Address,
		Courier:  shipmentOption.Name,
	}

	// convert to json
	jsonShippingData, _ := json.Marshal(&shippingData)
	jsonCartData, _ := json.Marshal(&cart)

	transaction := models.Transaction{
		UserID:         cart.UserID,
		StoreID:        uint(storeID),
		ShipmentStatus: "CHECKING", // checking, processing, delivered, cancelled
		PaymentStatus:  "PENDING",  // pending, paid, cancelled, refunded
		TrackingCode:   "",
		ShippingData:   string(jsonShippingData),
		CartData:       string(jsonCartData),
	}

	// insert transaction
	savedTransaction, err := database.InsertTransaction(&transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	detailTransaction := detailTransaction{}

	for _, item := range cart.CartDetails {
		// get product
		product, err := database.GetProductByID(int(item.ProductID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		transactionProduct := models.TransactionProduct{
			TransactionID: savedTransaction.ID.ID,
			Name:          product.Name,
			Slug:          product.Slug,
			Description:   product.Description,
			// Photo: product.,
			BasePrice: product.BasePrice,
			PriceCut:  product.PriceCut,
			OrderedAt: time.Now(),
		}

		// count reduced item
		product.Stock -= item.Qty
		_, errr := database.UpdateProductInfo(product)
		if errr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": errr.Error(),
			})
		}

		detailTransaction.transactionProducts = append(detailTransaction.transactionProducts, transactionProduct)
	}

	savedTransactionProducts, err := database.InsertTransactionProducts(&detailTransaction.transactionProducts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	savedTransaction.TransactionProducts = savedTransactionProducts

	// empty cart
	errr := database.EmptyCartByCartID(cartID)
	if errr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errr.Error(),
		})
	}

	cart.ShippingPrice = 0
	cart.Discount = 0
	cart.TotalPrice = 0

	_, errrr := database.UpdateCartInfo(cart)
	if errrr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errrr.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    savedTransaction,
	})
}
