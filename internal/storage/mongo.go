package storage

import (
	"Hospital_Service_API/internal/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage struct {
	db *mongo.Database
}

func NewMongoStorage(db *mongo.Database) *MongoStorage {
	return &MongoStorage{db: db}
}

func (s *MongoStorage) CreateAppointment(ctx context.Context, appointment models.Appointment) error {
	collection := s.db.Collection("appointments")

	// Check if the slot is already taken
	var existingAppointment models.Appointment
	filter := bson.M{"doctor_id": appointment.DoctorID, "slot": appointment.Slot}
	err := collection.FindOne(ctx, filter).Decode(&existingAppointment)
	if err == nil {
		return errors.New("slot is already taken")
	}
	if err != mongo.ErrNoDocuments {
		return err
	}

	_, err = collection.InsertOne(ctx, appointment)
	return err
}

func (s *MongoStorage) IsSlotAvailable(ctx context.Context, doctorID, slot string) bool {
	collection := s.db.Collection("appointments")
	doctorsCollection := s.db.Collection("doctors")

	// Парсим время слота
	slotTime, err := time.Parse(time.RFC3339, slot)
	if err != nil {
		return false // Невалидный формат времени
	}

	// Проверяем, что слот не является прошлым временем и не менее чем через час
	if slotTime.Before(time.Now()) || slotTime.Sub(time.Now()).Hours() < 1 {
		return false
	}

	// Проверяем, что слот существует у данного доктора
	doctorFilter := bson.M{"id": doctorID, "slots": slot}
	count, err := doctorsCollection.CountDocuments(ctx, doctorFilter)
	if err != nil || count == 0 {
		return false // Слот не найден у данного доктора
	}

	// Проверяем, что слот не занят
	appointmentFilter := bson.M{"doctor_id": doctorID, "slot": slot}
	count, err = collection.CountDocuments(ctx, appointmentFilter)
	return err == nil && count == 0
}

func (s *MongoStorage) GetAppointmentsBetween(ctx context.Context, start, end time.Time) ([]models.Appointment, error) {
	collection := s.db.Collection("appointments")
	filter := bson.M{
		"slot": bson.M{
			"$gte": start.Format(time.RFC3339),
			"$lte": end.Format(time.RFC3339),
		},
	}

	var appointments []models.Appointment
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var appointment models.Appointment
		if err := cursor.Decode(&appointment); err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, cursor.Err()
}

func (s *MongoStorage) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	collection := s.db.Collection("users")
	var user models.User
	filter := bson.M{"id": userID}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoStorage) GetDoctorByID(ctx context.Context, doctorID string) (*models.Doctor, error) {
	collection := s.db.Collection("doctors")
	var doctor models.Doctor
	filter := bson.M{"id": doctorID}
	if err := collection.FindOne(ctx, filter).Decode(&doctor); err != nil {
		return nil, err
	}
	return &doctor, nil
}
