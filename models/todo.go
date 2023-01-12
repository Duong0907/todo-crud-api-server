package todo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string             `json:"content,omitempty"`
	Status  string             `json:"status"`
}

// By default, the driver marshals the other fields to bson as the lowercase of the struct field name
