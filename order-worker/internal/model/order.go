package model

type Order struct {
	OrderId      string    `json:"order_id" bson:"order_id"`
	CustomerId   string    `json:"customer_id" bson:"customer_id"`
	CustomerName string    `json:"-" bson:"customer_name"`
	ProductIds   []string  `json:"product_ids" bson:"-"`
	Products     []Product `json:"-" bson:"products"`
}
