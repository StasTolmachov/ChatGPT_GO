package handlers

import (
	"ChatGPT_GO/user-service/auth"
	"ChatGPT_GO/user-service/database"
	"ChatGPT_GO/user-service/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// RegisterUser registerUser регистрирует нового пользователя
func RegisterUser(repo database.DbInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			log.Printf("c.ShouldBindJSON(&newUser): %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.DbPingMethod(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
			return
		}

		id, err := repo.RegisterUserMethod(newUser.Username, newUser.Password)

		if err != nil {
			log.Printf("database.Db.QueryRow(sqlInsert, newUser.Username, newUser.Password).Scan(&newUser.Id) %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		newUser.Id = id
		c.JSON(http.StatusCreated, newUser)
	}

}

func GetUserById(c *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))

	if database.Db == nil {
		fmt.Println("database.Db == nil")
		return
	}
	// Выполняем SQL-запрос
	err = database.Db.QueryRow("SELECT id, username FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username)
	if err != nil {
		// Логируем ошибку выполнения запроса
		log.Printf("Ошибка при запросе к базе данных: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка выполнения запроса"})
		return
	}

	log.Println("SQL-запрос успешно выполнен, обработка результата")

	// Возвращаем данные пользователя в формате JSON
	log.Println("Возвращаем данные пользователя в ответе")
	c.JSON(http.StatusOK, user)
}

// LoginHandler — обработчик логина
func LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Привязка данных JSON
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Проверка пользователя в базе данных
	var user models.User
	err := database.Db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", loginData.Username).
		Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		log.Printf("Ошибка запроса пользователя из базы данных: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Валидация пароля
	if loginData.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Генерация JWT токена
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		log.Printf("Ошибка генерации токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Возвращаем токен клиенту
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ProtectedRoute(c *gin.Context) {
	// Извлекаем имя пользователя из токена
	username := c.MustGet("username").(string)

	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + username + "!"})
}
