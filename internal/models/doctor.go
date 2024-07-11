package models

type Doctor struct {
	ID    string   `bson:"id" json:"id"`
	Name  string   `bson:"name" json:"name"`
	Spec  string   `bson:"spec" json:"spec"`
	Slots []string `bson:"slots" json:"slots"`
}
