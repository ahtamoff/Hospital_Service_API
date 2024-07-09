package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Doctor struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Spec  string             `bson:"spec" json:"spec"`
	Slots []time.Time        `bson:"slots" json:"slots"`
}
