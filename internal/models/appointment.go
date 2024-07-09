package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Appointment struct {
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	DoctorID primitive.ObjectID `bson:"doctor_id" json:"doctor_id"`
	Slot     time.Time          `bson:"slot" json:"slot"`
}
