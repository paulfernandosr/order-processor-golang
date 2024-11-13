package model

type Product struct {
	ProductId   string  `json:"product_id" bson:"product_id"`
	Name        string  `json:"name" bson:"name" binding:"required"`
	Description string  `json:"description" bson:"description" binding:"required"`
	Price       float64 `json:"price" bson:"price" binding:"required"`
}
