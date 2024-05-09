package models

type Product struct {
	UserId      string `json:"id"`
	Name        string `json:"name"`
	SKU         string `json:"sku"`
	Category    string `json:"category"`
	ImagerUrl   string `json:"imageUrl"`
	Notes       string `json:"notes"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Location    string `json:"location"`
	IsAvailable bool   `json:"isAvailable"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}