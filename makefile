# Путь к каждому микросервису
USER_SERVICE_PATH=./user-service
TASK_SERVICE_PATH=./task-service
API_GATEWAY_PATH=./api-gateway

# Команда для запуска каждого сервиса
.PHONY: run

run:
	@echo "Запуск User Service..."
	cd $(USER_SERVICE_PATH) && go run main.go &

	@echo "Запуск Task Service..."
	cd $(TASK_SERVICE_PATH) && go run main.go &

	@echo "Запуск API Gateway..."
	cd $(API_GATEWAY_PATH) && go run main.go &

	@echo "Все микросервисы запущены"

