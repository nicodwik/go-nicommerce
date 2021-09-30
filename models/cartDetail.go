package models

type CartDetail struct {
	ID
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Qty       int  `json:"qty"`
	Timestamp
}
