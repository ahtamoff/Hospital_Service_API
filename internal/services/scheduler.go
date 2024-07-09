package services

import (
	"time"

	"github.com/ahtamoff/Hospital_Service_API/internal/storage"
)

type Scheduler struct {
	Storage             *storage.Storage
	NotificationService *NotificationService
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			s.checkAppointments()
		}
	}()
}

func (s *Scheduler) checkAppointments() {
	appointments, _ := s.Storage.GetUpcomingAppointments()

	for _, appointment := range appointments {
		user, _ := s.Storage.GetUser(appointment.UserID)
		doctor, _ := s.Storage.GetDoctor(appointment.DoctorID)
		s.NotificationService.Notify(appointment, user, doctor, time.Now())
	}
}
