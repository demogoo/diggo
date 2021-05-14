package models

type Order struct {
	ID    int `json:"order_id"    validate:"required"`
	Price int `json:"price" validate:"required"`
}
