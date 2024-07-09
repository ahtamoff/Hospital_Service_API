package storage

import (
	"context"
	"time"

	"Hospital_Service_API/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client *mongo.Client
}

func NewStorage(uri string) (*Storage, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &Storage{Client: client}, nil
}

func (s *Storage) GetDoctor(id primitive.ObjectID) (models.Doctor, error) {
	var doctor models.Doctor
	collection := s.Client.Database("appointment_service").Collection("doctors")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&doctor)
	return doctor, err
}

func (s *Storage) GetUser(id primitive.ObjectID) (models.User, error) {
	var user models.User
	collection := s.Client.Database("appointment_service").Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (s *Storage) BookSlot(appointment models.Appointment) error {
	collection := s.Client.Database("appointment_service").Collection("appointments")
	_, err := collection.InsertOne(context.Background(), appointment)
	return err
}

func (s *Storage) GetUpcomingAppointments() ([]models.Appointment, error) {
	var appointments []models.Appointment
	collection := s.Client.Database("appointment_service").Collection("appointments")
	cursor, err := collection.Find(context.Background(), bson.M{
		"slot": bson.M{
			"$gte": time.Now(),
		},
	})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &appointments); err != nil {
		return nil, err
	}
	return appointments, nil
}
