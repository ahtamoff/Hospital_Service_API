package main

import (
	"Hospital_Service_API/config"
	"Hospital_Service_API/internal/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.LoadConfig()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database(cfg.DBName)

	// Заполнение данных пользователей
	users := []interface{}{
		models.User{ID: "user1", Phone: "+7 922 548 54 23", Name: "Василий"},
		models.User{ID: "user2", Phone: "+7 934 765 45 69", Name: "Биба"},
		models.User{ID: "user3", Phone: "+7 987 345 56 32", Name: "Боба"},
	}
	_, err = db.Collection("users").InsertMany(context.Background(), users)
	if err != nil {
		log.Fatal(err)
	}

	// Заполнение данных врачей
	doctors := []interface{}{
		models.Doctor{ID: "doctor1", Name: "Светлана", Spec: "Терапевт", Slots: []string{
			time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			time.Now().Add(48 * time.Hour).Format(time.RFC3339),
		}},
		models.Doctor{ID: "doctor2", Name: "Алексей", Spec: "Окулист", Slots: []string{
			time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			time.Now().Add(48 * time.Hour).Format(time.RFC3339),
		}},
	}
	_, err = db.Collection("doctors").InsertMany(context.Background(), doctors)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Данные успешно загружены в базу данных")
}
