package models

import "time"

type TransactionProduct struct {
	ID
	TransactionID uint      `json:"transaction_id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	Photo         string    `json:"photo"`
	BasePrice     int       `json:"base_price"`
	PriceCut      int       `json:"price_cut"`
	OrderedAt     time.Time `json:"ordered_at"`
	Timestamp
}
