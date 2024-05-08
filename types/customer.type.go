package types

type CustomerRegisterRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}