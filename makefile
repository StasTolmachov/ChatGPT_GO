# Путь к каждому микросервису
USER_SERVICE_PATH=./user-service
TASK_SERVICE_PATH=./task-service
API_GATEWAY_PATH=./api-gateway

# Файлы для сохранения PID
USER_SERVICE_PID=user_service.pid
TASK_SERVICE_PID=task_service.pid
API_GATEWAY_PID=api_gateway.pid

# Команда для запуска каждого сервиса
.PHONY: run stop

run:
	@echo "Запуск User Service..."
	cd $(USER_SERVICE_PATH) && go run main.go & echo $$! > $(USER_SERVICE_PID)

	@echo "Запуск Task Service..."
	cd $(TASK_SERVICE_PATH) && go run main.go & echo $$! > $(TASK_SERVICE_PID)

	@echo "Запуск API Gateway..."
	cd $(API_GATEWAY_PATH) && go run main.go & echo $$! > $(API_GATEWAY_PID)

	@echo "Все микросервисы запущены"

# Команда для остановки каждого сервиса
stop:
	@echo "Остановка User Service..."
	@kill `cat $(USER_SERVICE_PID)` && rm $(USER_SERVICE_PID)

	@echo "Остановка Task Service..."
	@kill `cat $(TASK_SERVICE_PID)` && rm $(TASK_SERVICE_PID)

	@echo "Остановка API Gateway..."
	@kill `cat $(API_GATEWAY_PID)` && rm $(API_GATEWAY_PID)

	@echo "Все микросервисы остановлены"
