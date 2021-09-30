package models

type ProductGallery struct {
	ID
	ProductID uint   `json:"product_id"`
	Photo     string `json:"photo"`
	IsPrimary bool   `json:"is_primary"`
	Timestamp
}
