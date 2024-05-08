package models

type Customer struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}