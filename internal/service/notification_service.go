package service

import (
	"Hospital_Service_API/internal/storage"
	"context"
	"log"
	"os"
	"time"
)

type NotificationService struct {
	storage *storage.MongoStorage
}

func NewNotificationService(storage *storage.MongoStorage) *NotificationService {
	return &NotificationService{storage: storage}
}

func (s *NotificationService) StartNotificationScheduler() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	notifiedAppointmentsTomorrow := make(map[string]bool)
	notifiedAppointmentsTwoHour := make(map[string]bool)
	for {
		s.sendNotifications(notifiedAppointmentsTomorrow, notifiedAppointmentsTwoHour)
		<-ticker.C
	}
}

func (s *NotificationService) sendNotifications(notifiedAppointmentsTomorrow, notifiedAppointmentsTwoHour map[string]bool) {
	logFile, err := os.OpenFile("notifications.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	ctx := context.Background()
	now := time.Now()
	// Определяем завтрашний день с начала до конца
	tomorrowStart := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	tomorrowEnd := tomorrowStart.Add(24 * time.Hour)

	// Два часа от текущего времени
	twoHoursLaterStart := now.Add(2 * time.Hour)
	twoHoursLaterEnd := twoHoursLaterStart.Add(1 * time.Hour)

	// Получаем записи на завтра
	appointmentsTomorrow, err := s.storage.GetAppointmentsBetween(ctx, tomorrowStart, tomorrowEnd)
	if err != nil {
		logger.Println("Error fetching appointments for tomorrow:", err)
		return
	}

	// Получаем записи через два часа
	appointmentsTwoHoursLater, err := s.storage.GetAppointmentsBetween(ctx, twoHoursLaterStart, twoHoursLaterEnd)
	if err != nil {
		logger.Println("Error fetching appointments for two hours later:", err)
		return
	}

	for _, appointment := range appointmentsTomorrow {
		user, err := s.storage.GetUserByID(ctx, appointment.UserID)
		if err != nil {
			logger.Println("Error fetching user:", err)
			continue
		}
		doctor, err := s.storage.GetDoctorByID(ctx, appointment.DoctorID)
		if err != nil {
			logger.Println("Error fetching doctor:", err)
			continue
		}
		if _, ok := notifiedAppointmentsTomorrow[appointment.ID]; !ok {
			logger.Printf("%s | Привет %s! Напоминаем что вы записаны к %s завтра в %s!", now.Format(time.RFC3339), user.Name, doctor.Spec, appointment.Slot)
			notifiedAppointmentsTomorrow[appointment.ID] = true
		}
	}

	for _, appointment := range appointmentsTwoHoursLater {
		user, err := s.storage.GetUserByID(ctx, appointment.UserID)
		if err != nil {
			logger.Println("Error fetching user:", err)
			continue
		}
		doctor, err := s.storage.GetDoctorByID(ctx, appointment.DoctorID)
		if err != nil {
			logger.Println("Error fetching doctor:", err)
			continue
		}
		if _, ok := notifiedAppointmentsTwoHour[appointment.ID]; !ok {
			logger.Printf("%s | Привет %s! Вам через 2 часа к %s в %s!", now.Format(time.RFC3339), user.Name, doctor.Spec, appointment.Slot)
			notifiedAppointmentsTwoHour[appointment.ID] = true
		}
	}
}
