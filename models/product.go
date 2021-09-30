package models

type Product struct {
	ID
	CategoryID  uint   `json:"category_id"`
	StoreID     uint   `json:"store_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Weight      int    `json:"weight"`
	BasePrice   int    `json:"base_price"`
	PriceCut    int    `json:"price_cut"`

	// Relationship
	ProductGalleries []ProductGallery `json:"product_galleries"`
	// Transactions     []Transaction    `json:"transactions"`
	Timestamp
}
