package model

type Customer struct {
	CustomerId string `bson:"customer_id" json:"customer_id" binding:"required"`
	Name       string `json:"name" bson:"name" binding:"required"`
	Email      string `json:"email" bson:"email" binding:"required"`
	Address    string `json:"address" bson:"address"`
	IsActive   bool   `json:"is_active" bson:"is_active" binding:"required"`
}
