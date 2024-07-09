package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Phone string             `bson:"phone" json:"phone"`
	Name  string             `bson:"name" json:"name"`
}
