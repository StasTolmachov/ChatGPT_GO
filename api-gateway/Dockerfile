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
COPY . .

# Переходим в директорию task-service для сборки
WORKDIR /app/api-gateway

# Собираем приложение
RUN go build -o api-gateway .

# Открываем порт
EXPOSE 8082

# Запускаем приложение
CMD ["./api-gateway"]
