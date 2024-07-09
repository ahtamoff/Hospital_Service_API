package main

import (
	"log"
	"net"

	"github.com/ahtamoff/Hospital_Service_API/internal/services"
	"github.com/ahtamoff/Hospital_Service_API/internal/storage"
	pb "github.com/ahtamoff/Hospital_Service_API/pb"
	"google.golang.org/grpc"
)

func main() {
	storage, err := storage.NewStorage("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	appointmentService := &services.AppointmentService{Storage: storage}
	notificationService := &services.NotificationService{}
	scheduler := &services.Scheduler{Storage: storage, NotificationService: notificationService}

	grpcServer := grpc.NewServer()
	pb.RegisterAppointmentServiceServer(grpcServer, &grpc.Server{AppointmentService: appointmentService})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	go scheduler.Start()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
