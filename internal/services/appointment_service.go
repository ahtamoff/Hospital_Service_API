package services

import (
	"errors"
	"github.com/ahtamoff/Hospital_Service_API/internal/models"
	"github.com/ahtamoff/Hospital_Service_API/internal/storage"
)

type AppointmentService struct {
	Storage *storage.Storage
}

func (s *AppointmentService) BookAppointment(appointment models.Appointment) error {
	doctor, err := s.Storage.GetDoctor(appointment.DoctorID)
	if err != nil {
		return err
	}

	slotAvailable := false
	for _, slot := range doctor.Slots {
		if slot.Equal(appointment.Slot) {
			slotAvailable = true
			break
		}
	}

	if !slotAvailable {
		return errors.New("slot not available")
	}

	return s.Storage.BookSlot(appointment)
}
