# Используем официальный образ Go
FROM golang:1.24

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum до копирования остального кода
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник
RUN go build -o main ./cmd/main.go

# Указываем порт
EXPOSE 8080

# Запуск приложения
CMD ["./main"]
