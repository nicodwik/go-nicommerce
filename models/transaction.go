package models

type Transaction struct {
	ID
	UserID         uint   `json:"user_id"`
	StoreID        uint   `json:"store_id"`
	ShipmentStatus string `json:"shipment_status"`
	PaymentStatus  string `json:"payment_status"`
	TrackingCode   string `json:"tracking_code"`
	ShippingData   string `json:"shipping_data"`
	CartData       string `json:"cart_data"`

	// Relationship
	TransactionProducts []TransactionProduct `json:"transaction_products"`
	Timestamp
}
