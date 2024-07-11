# Hospital_Service_API
API сервиса по записи к врачам

# Запуск сервера:
    docker-compose up -d
# Выполнить миграцию:
    docker-compose run --rm app go run migrations/main.go
# Правильно записаться:
curl -X POST http://localhost:8080/appointments -d '{
"user_id": "user1",    
"doctor_id": "doctor2",    
"slot": "2024-07-12T17:20:24Z"                    
}' -H "Content-Type: application/json"

# Неправильно записаться:
curl -X POST http://localhost:8080/appointments -d '{
"user_id": "user1",    
"doctor_id": "doctor2",    
"slot": "2019-07-12T17:20:24Z"                    
}' -H "Content-Type: application/json"
(или любой несуществующий слот)
