package models

type Category struct {
	ID
	StoreID uint   `json:"store_id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`

	// Relationship
	Products []Product `json:"products"`
	Timestamp
}
