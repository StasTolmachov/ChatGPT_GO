# Используем официальный образ Golang
FROM golang:1.23

# Создаем рабочую директорию для проекта
WORKDIR /app

# Копируем go.mod и go.sum из корневой директории проекта
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем только директорию с исходниками task-service
COPY ./task-service ./task-service

# Переходим в директорию task-service для сборки
WORKDIR /app/task-service

# Собираем приложение
RUN go build -o task-service .

# Открываем порт
EXPOSE 8082

# Запускаем приложение
CMD ["./task-service"]
