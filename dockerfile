# этап сборки
FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/main.go

# этап рантайма
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
# **Добавляем** папку миграций
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations
# и остальное
COPY init.sql ./
COPY .env ./

RUN apk add --no-cache tzdata

EXPOSE 8080
CMD ["./main"]
