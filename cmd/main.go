package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Hospital_Service_API/config"
	"Hospital_Service_API/internal/handlers"
	"Hospital_Service_API/internal/service"
	"Hospital_Service_API/internal/storage"
)

func main() {
	cfg := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database(cfg.DBName)
	mongoStorage := storage.NewMongoStorage(db)

	// Создание роутера Gin
	router := gin.Default()
	handlers.SetupRoutes(router, mongoStorage)

	// Запуск сервиса оповещений
	notificationService := service.NewNotificationService(mongoStorage)
	go notificationService.Start()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
