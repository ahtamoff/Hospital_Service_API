# Hospital_Service_API
API сервиса по записи к врачам

# Запуск сервера:
    make all
# Инициализировать базу данных:
    make init_db


# Проверка функционала через Postman (Для удобной проверки частота записи напоминаний 1 раз / 5 сек)
        POST 
        http://localhost:8080/appointments
        body(заменить время слота на актуальное из бд):
        {
        "user_id": "user3",
        "doctor_id": "doctor1",
        "slot": "2024-07-13T09:06:37Z" 
        }
        
