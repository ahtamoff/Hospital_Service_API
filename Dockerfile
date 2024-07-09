FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /go-appointment-service cmd/main.go

EXPOSE 50051

CMD ["/go-appointment-service"]
