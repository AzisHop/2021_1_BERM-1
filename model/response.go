package model

type Response struct {
	ID        uint64 `json:"id,omitempty" db:"id"`
	OrderID   uint64 `json:"order_id" db:"order_id"`
	UserID    uint64 `json:"user_id" db:"user_id"`
	Rate      uint64 `json:"rate" db:"rate"`
	UserLogin string `json:"user_login" db:"user_login"`
	UserImg   string `json:"user_img" db:"user_img"`
}