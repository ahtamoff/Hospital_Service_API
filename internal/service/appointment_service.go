package service

import (
	"Hospital_Service_API/internal/models"
	"Hospital_Service_API/internal/storage"
	"context"
	"errors"
)

type AppointmentService struct {
	storage *storage.MongoStorage
}

func NewAppointmentService(storage *storage.MongoStorage) *AppointmentService {
	return &AppointmentService{storage: storage}
}

func (s *AppointmentService) CreateAppointment(ctx context.Context, userID, doctorID, slot string) error {
	appointment := models.Appointment{
		UserID:   userID,
		DoctorID: doctorID,
		Slot:     slot,
	}

	// Check if the slot is available
	if !s.storage.IsSlotAvailable(ctx, doctorID, slot) {
		return errors.New("slot is not available")
	}

	// Create the appointment
	if err := s.storage.CreateAppointment(ctx, appointment); err != nil {
		return err
	}

	return nil
}
