# Используем официальный образ Go как базовый
FROM golang:1.19-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исходники приложения в рабочую директорию
COPY . .

# Инициализируем модуль и скачиваем зависимости
RUN go mod tidy

# Собираем приложение
RUN go build -o main

# Минимизируем конечный образ
FROM alpine:latest
RUN apk add --no-cache curl
COPY --from=builder /app/main /main
EXPOSE 8080
CMD ["/main"]
