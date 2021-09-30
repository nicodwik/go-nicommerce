package models

type Cart struct {
	ID
	UserID        uint `json:"user_id"`
	ShippingPrice int  `json:"shipping_price"`
	Discount      int  `json:"discount"`
	TotalPrice    int  `json:"total_price"`

	//Relationship
	CartDetails []CartDetail
	Timestamp
}
