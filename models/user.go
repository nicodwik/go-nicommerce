package models

type User struct {
	ID
	Firstname   string `json:"firstname" form:"firstname"`
	Lastname    string `json:"lastname" form:"lastname"`
	Email       string `json:"email" form:"email"`
	Phone       string `json:"phone" form:"phone"`
	Avatar      string `json:"avatar" form:"avatar"`
	Password    string `json:"-" form:"password"`
	Token       string `json:"token"`
	StoreStatus bool   `json:"store_status" form:"store_status"`

	// Relationship
	AddressOptions []AddressOption `json:"address_options"`
	Store          Store           `json:"store"`
	Cart           Cart            `json:"cart"`
	Timestamp                      // created_at, updated_at, deleted_at
}
