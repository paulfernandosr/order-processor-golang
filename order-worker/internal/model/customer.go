package model

type Customer struct {
	CustomerId string `json:"customer_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	IsActive   bool   `json:"is_active"`
}
