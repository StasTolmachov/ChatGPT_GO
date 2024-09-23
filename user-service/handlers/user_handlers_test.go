package handlers

import (
	"ChatGPT_GO/user-service/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Тест для функции RegisterUser
func TestRegisterUser(t *testing.T) {
	// Создаем мок-репозиторий
	mockRepo := new(database.MockUserRepository)

	// Задаем поведение мока
	mockRepo.On("DbPingMethod").Return(nil)                                 // Соединение с базой успешно
	mockRepo.On("RegisterUserMethod", "john", "password123").Return(1, nil) // Новый пользователь с ID 1

	// Создаем запрос
	req, _ := http.NewRequest("POST", "/users/register", strings.NewReader(`{"Username": "john", "Password": "password123"}`))
	req.Header.Set("Content-Type", "application/json")

	// Создаем объект для записи ответа
	w := httptest.NewRecorder()

	// Создаем контекст Gin
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Вызываем обработчик с мок-репозиторием
	handler := RegisterUser(mockRepo)
	handler(c)

	// Проверяем ответ
	assert.Equal(t, http.StatusCreated, w.Code)

	// Проверяем, что JSON содержит ID пользователя
	expected := `{"Id":1,"Username":"john","Password":"password123"}`
	assert.JSONEq(t, expected, w.Body.String())

	// Проверяем, что методы Ping и CreateUser были вызваны
	mockRepo.AssertExpectations(t)
}
