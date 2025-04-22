# Gunakan versi Go yang sesuai dengan go.mod
FROM golang:1.23.4 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o main .

CMD ["/app/main"]