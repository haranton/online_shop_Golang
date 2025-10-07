# === Этап 1: сборка и тестирование ===
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Скопировать модули и скачать зависимости
COPY go.mod go.sum ./
RUN go mod download

# Скопировать исходники
COPY . .

# --- Запуск юнит-тестов (прерывает сборку при ошибках) ---
RUN go test -v ./...

# --- Сборка бинарника ---
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/main.go


# === Этап 2: продакшн ===
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations
COPY init.sql ./
COPY .env ./

RUN apk add --no-cache tzdata

EXPOSE 8080
CMD ["./main"]
