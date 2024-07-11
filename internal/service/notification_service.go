package service

import (
	"Hospital_Service_API/internal/storage"
	"context"
	"log"
	"time"
)

type NotificationService struct {
	storage *storage.MongoStorage
}

func NewNotificationService(storage *storage.MongoStorage) *NotificationService {
	return &NotificationService{storage: storage}
}

func (s *NotificationService) Start() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.sendNotifications()
			}
		}
	}()
}

func (s *NotificationService) sendNotifications() {
	ctx := context.Background()
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	twoHoursLater := now.Add(2 * time.Hour)

	appointmentsTomorrow, err := s.storage.GetAppointmentsBetween(ctx, tomorrow, tomorrow.Add(time.Hour))
	if err != nil {
		log.Println("Error fetching appointments for tomorrow:", err)
		return
	}

	appointmentsTwoHoursLater, err := s.storage.GetAppointmentsBetween(ctx, twoHoursLater, twoHoursLater.Add(time.Hour))
	if err != nil {
		log.Println("Error fetching appointments for two hours later:", err)
		return
	}

	for _, appointment := range appointmentsTomorrow {
		user, err := s.storage.GetUserByID(ctx, appointment.UserID)
		if err != nil {
			log.Println("Error fetching user:", err)
			continue
		}
		doctor, err := s.storage.GetDoctorByID(ctx, appointment.DoctorID)
		if err != nil {
			log.Println("Error fetching doctor:", err)
			continue
		}
		log.Printf("%s | Привет %s! Напоминаем что вы записаны к %s завтра в %s!", now.Format(time.RFC3339), user.Name, doctor.Spec, appointment.Slot)
	}

	for _, appointment := range appointmentsTwoHoursLater {
		user, err := s.storage.GetUserByID(ctx, appointment.UserID)
		if err != nil {
			log.Println("Error fetching user:", err)
			continue
		}
		doctor, err := s.storage.GetDoctorByID(ctx, appointment.DoctorID)
		if err != nil {
			log.Println("Error fetching doctor:", err)
			continue
		}
		log.Printf("%s | Привет %s! Вам через 2 часа к %s в %s!", now.Format(time.RFC3339), user.Name, doctor.Spec, appointment.Slot)
	}
}
