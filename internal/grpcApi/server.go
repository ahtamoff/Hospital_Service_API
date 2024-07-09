package grpcApi

import (
	"context"
	"github.com/ahtamoff/Hospital_Service_API/internal/models"
	"github.com/ahtamoff/Hospital_Service_API/internal/services"
	pb "github.com/ahtamoff/Hospital_Service_API/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	pb.UnimplementedAppointmentServiceServer
	AppointmentService *services.AppointmentService
}

func (s *Server) BookAppointment(ctx context.Context, req *pb.BookAppointmentRequest) (*pb.BookAppointmentResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return &pb.BookAppointmentResponse{Status: "error", Message: "Invalid user ID"}, nil
	}

	doctorID, err := primitive.ObjectIDFromHex(req.DoctorId)
	if err != nil {
		return &pb.BookAppointmentResponse{Status: "error", Message: "Invalid doctor ID"}, nil
	}

	slot := req.Slot.AsTime()
	appointment := models.Appointment{
		UserID:   userID,
		DoctorID: doctorID,
		Slot:     slot,
	}

	if err := s.AppointmentService.BookAppointment(appointment); err != nil {
		return &pb.BookAppointmentResponse{Status: "error", Message: err.Error()}, nil
	}

	return &pb.BookAppointmentResponse{Status: "success", Message: "Appointment booked successfully"}, nil
}
