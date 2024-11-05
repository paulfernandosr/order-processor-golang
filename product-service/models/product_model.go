package models

type Product struct {
	ProductId   string  `json:"product_id" bson:"product_id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
}
