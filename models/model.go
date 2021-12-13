package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Watched bool               `json:"watched"`
}
