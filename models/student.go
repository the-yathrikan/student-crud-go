package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Age    int                `json:"age" bson:"age"`
	Course string             `json:"course" bson:"course"`
}
