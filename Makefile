# Имя контейнера приложения
APP_CONTAINER_NAME := hospital_service_api-app-1

# Имя контейнера MongoDB
MONGO_CONTAINER_NAME := hospital_service_api-mongo-1

# Файлы и директории
MIGRATIONS_DIR := migrations
MIGRATION_SCRIPT := $(MIGRATIONS_DIR)/main.go
DOCKER_COMPOSE_FILE := docker-compose.yml

# Команды для Makefile
.PHONY: all build up down migrate init_db logs

# Собрать и запустить все контейнеры
all: build up

# Собрать контейнеры
build:
	@docker-compose build

# Запустить контейнеры в фоновом режиме
up:
	@docker-compose up -d

# Остановить контейнеры
down:
	@docker-compose down

# Выполнить миграции (инициализация базы данных)
migrate:
	@docker-compose run --rm app go run $(MIGRATION_SCRIPT)

# Инициализировать базу данных
init_db: migrate


# Показать логи контейнера приложения
logs:
	@docker-compose logs app
