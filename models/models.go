package models

import "time"

type User struct {
	ID
	Firstname   string `json:"firstname" form:"firstname"`
	Lastname    string `json:"lastname" form:"lastname"`
	Email       string `json:"email" form:"email"`
	Phone       string `json:"phone" form:"phone"`
	Avatar      string `json:"avatar" form:"avatar"`
	Password    string `json:"-" form:"password"`
	StoreStatus bool   `json:"store_status" form:"store_status"`

	// Relationship
	AddressOptions []AddressOption `json:"address_options"`
	Store          Store           `json:"store"`
	Cart           Cart            `json:"cart"`
	Timestamp                      // created_at, updated_at, deleted_at
}

type AddressOption struct {
	ID
	UserID     uint   `json:"user_id"`
	ProvinceID int    `json:"province_id"`
	CityID     int    `json:"city_id"`
	Address    string `json:"address"`
	Timestamp
}

type Cart struct {
	ID
	UserID        uint `json:"user_id"`
	ShippingPrice int  `json:"shipping_price"`
	Tax           int  `json:"tax"`
	Discount      int  `json:"discount"`
	TotalPrice    int  `json:"total_price"`

	//Relationship
	CartDetails []CartDetail
	Timestamp
}

type CartDetail struct {
	ID
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Qty       int  `json:"qty"`
	Timestamp
}

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

type ShipmentOption struct {
	ID
	StoreID  uint   `json:"store_id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Avatar   string `json:"avatar"`
	Timestamp
}

type Category struct {
	ID
	StoreID uint   `json:"store_id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`

	// Relationship
	Products []Product `json:"products"`
	Timestamp
}

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

type ProductGallery struct {
	ID
	ProductID uint `json:"product_id"`
	IsPrimary bool `json:"is_primary"`
	Timestamp
}

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
