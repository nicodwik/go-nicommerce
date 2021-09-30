package models

type Store struct {
	ID
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	ProvinceID  int    `json:"province_id"`
	CityID      int    `json:"city_id"`
	Address     string `json:"address"`

	// Relationship
	ShipmentOptions []ShipmentOption `json:"shipment_options"`
	Categories      []Category       `json:"categories"`
	Transactions    []Transaction    `json:"transactions"`
	// Products        []Product        `json:"products"`

	Timestamp
}
