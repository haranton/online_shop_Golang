# === Этап 1: сборка приложения ===
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Копируем модули и зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник (без CGO для простоты)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/main.go


# === Этап 2: продакшн-образ ===
FROM alpine:latest

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/main .

# Копируем миграции, SQL и env
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations
COPY .env ./

# Добавляем временную зону (для логов и т.п.)
RUN apk add --no-cache tzdata

EXPOSE 8080
CMD ["./main"]
