package services

import (
	"log"
	"time"

	"github.com/ahtamoff/Hospital_Service_API/internal/models"
)

type NotificationService struct{}

func (n *NotificationService) Notify(appointment models.Appointment, user models.User, doctor models.Doctor, current time.Time) {
	if current.Add(24 * time.Hour).After(appointment.Slot) {
		log.Printf("%s | Привет %s! Напоминаем что вы записаны к %s завтра в %s!\n", current.Format(time.RFC3339), user.Name, doctor.Spec, appointment.Slot.Format(time.RFC3339))
	}

	if current.Add(2 * time.Hour).After(appointment.Slot) {
		log.Printf("%s | Привет %s! Вам через 2 часа к %s в %s!\n", current.Format(time.RFC3339), user.Name, doctor.Spec, appointment.Slot.Format(time.RFC3339))
	}
}
