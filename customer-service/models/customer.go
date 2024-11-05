package models

type Customer struct {
	CustomerId string `bson:"customer_id" json:"customer_id"`
	Name       string `json:"name" bson:"name"`
	Email      string `json:"email" bson:"email"`
	Address    string `json:"address" bson:"address"`
}
