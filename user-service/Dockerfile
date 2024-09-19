# Используем официальный образ Golang
FROM golang:1.23

# Создаем рабочую директорию для проекта
WORKDIR /app

# Копируем go.mod и go.sum из корневой директории проекта
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем только директорию с исходниками user-service
COPY ./user-service ./user-service

# Переходим в директорию user-service для сборки
WORKDIR /app/user-service

# Собираем приложение
RUN go build -o user-service .

# Открываем порт
EXPOSE 8081

# Запускаем приложение
CMD ["./user-service"]
