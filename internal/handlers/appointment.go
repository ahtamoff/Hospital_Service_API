package handlers

import (
	"Hospital_Service_API/internal/service"
	"Hospital_Service_API/internal/storage"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppointmentRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	DoctorID string `json:"doctor_id" binding:"required"`
	Slot     string `json:"slot" binding:"required"`
}

func SetupRoutes(router *gin.Engine, storage *storage.MongoStorage) {
	appointmentService := service.NewAppointmentService(storage)
	router.POST("/appointments", createAppointmentHandler(appointmentService))
}

func createAppointmentHandler(service *service.AppointmentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AppointmentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := context.Background()
		err := service.CreateAppointment(ctx, req.UserID, req.DoctorID, req.Slot)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "appointment created"})
	}
}
