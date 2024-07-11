package models

type User struct {
	ID    string `bson:"id" json:"id"`
	Phone string `bson:"phone" json:"phone"`
	Name  string `bson:"name" json:"name"`
}
