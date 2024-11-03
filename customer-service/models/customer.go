package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}
