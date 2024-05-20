package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-nicommerce/env"
	"go-nicommerce/lib/database"
	"go-nicommerce/middlewares"
	"go-nicommerce/models"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/random"
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
	cartID := middlewares.ExtractTokenUserId(c)
	_, err := database.GetCartByID(cartID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Cart Not Found",
		})
	}

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

	now := time.Now()
	dateNow := now.Format("02012006")
	randomString := strings.ToUpper(random.String(8))

	transaction := models.Transaction{
		UserID:         cart.UserID,
		StoreID:        uint(storeID),
		OrderCode:      fmt.Sprintf("TRX-%v-%v", dateNow, randomString),
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

	//process to payment gateway
	transactionDetail := transactionDetail{
		OrderId:     savedTransaction.OrderCode,
		GrossAmount: cart.TotalPrice,
	}

	customerDetail := customerDetail{
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	paymentGatewayPayload := paymentGatewayPayload{
		TransactionDetails: transactionDetail,
		CustomerDetail:     customerDetail,
	}

	paymentUrl, err := _processToPaymentGateway(paymentGatewayPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

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
		"data": map[string]interface{}{
			"transaction": savedTransaction,
			"payment": map[string]string{
				"redirect_url": paymentUrl,
			},
		},
	})
}

type paymentGatewayPayload struct {
	TransactionDetails transactionDetail `json:"transaction_details"`
	CustomerDetail     customerDetail    `json:"customer_details"`
}

type transactionDetail struct {
	OrderId     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

type customerDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func _processToPaymentGateway(payload paymentGatewayPayload) (string, error) {
	serverKey := env.Find("MIDTRANS_SERVER_KEY", "SB-Mid-server-fS1mtLoitm8Pb3DB3X2Hb5pO:")
	authKey := base64.StdEncoding.EncodeToString([]byte(serverKey))

	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"

	jsonPayload, _ := json.Marshal(payload)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	request.Header.Set("accept", "application/json")
	request.Header.Set("content-type", "application/json")
	request.Header.Set("authorization", fmt.Sprintf("basic %v", authKey))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	type resp struct {
		Token       string `json:"token"`
		RedirectUrl string `json:"redirect_url"`
	}

	r := resp{}

	body, _ := io.ReadAll(response.Body)

	_ = json.Unmarshal(body, &r)
	return r.RedirectUrl, nil
}

func PaymentCallback(c echo.Context) error {
	reqData := make(map[string]interface{})
	_ = json.NewDecoder(c.Request().Body).Decode(&reqData)

	// get order_id from request body
	orderCode, ok := reqData["order_id"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "bad request",
		})
	}

	transaction, err := database.GetTransactionByOrderCode(orderCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if transaction.PaymentStatus != "PENDING" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "payment already processed (not in pending status)",
		})
	}

	// get transaction_status from request body
	transactionStatus, ok := reqData["transaction_status"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "bad request",
		})
	}

	// set payment status from transaction_status
	switch transactionStatus {
	case "settlement":
		transaction.PaymentStatus = "PAID"

	case "pending":
		transaction.PaymentStatus = "PENDING"

	default:
		transaction.PaymentStatus = "CANCEL"

		//restore qty each product
		trxCartData := make(map[string]interface{})
		_ = json.Unmarshal([]byte(transaction.CartData), &trxCartData)

		s := reflect.ValueOf(trxCartData["CartDetails"])

		cartDetails := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			cartDetails[i] = s.Index(i).Interface()
		}

		for _, cd := range cartDetails {
			d, _ := cd.(map[string]interface{})

			productId := int(d["product_id"].(float64))
			qty := int(d["qty"].(float64))

			product, err := database.GetProductByID(productId)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": err.Error(),
				})
			}

			product.Stock += qty
			_, err = database.UpdateProductInfo(product)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": err.Error(),
				})
			}
		}
	}

	savedTransaction, err := database.UpdateTransactionData(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": map[string]interface{}{
			"order_code": savedTransaction.OrderCode,
		},
	})
}
