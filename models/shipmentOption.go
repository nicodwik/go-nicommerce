package models

type ShipmentOption struct {
	ID
	StoreID  uint   `json:"store_id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Avatar   string `json:"avatar"`
	Timestamp
}
