FROM golang:1.22.4-alpine

WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Сборка приложения
RUN go build -o /appointment-service ./cmd

EXPOSE 8080

CMD [ "/appointment-service" ]
