package models

type AddressOption struct {
	ID
	UserID     uint   `json:"user_id"`
	ProvinceID int    `json:"province_id"`
	CityID     int    `json:"city_id"`
	Address    string `json:"address"`
	Timestamp
}
