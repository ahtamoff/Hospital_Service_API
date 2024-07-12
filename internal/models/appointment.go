package models

type Appointment struct {
	ID       string `bson:"id" json:"id"`
	UserID   string `bson:"user_id" json:"user_id"`
	DoctorID string `bson:"doctor_id" json:"doctor_id"`
	Slot     string `bson:"slot" json:"slot"`
}
