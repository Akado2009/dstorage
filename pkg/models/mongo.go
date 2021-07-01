package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Info is a file info across servers
type Info struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Servers   []string           `bson:"servers,omitempty"`
	CheckSums []string           `bson:"checkSums,omitempty"`
}
