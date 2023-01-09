package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Config struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Value       string             `json:"value,omitempty"`
	Description string             `json:"description,omitempty"`
}
