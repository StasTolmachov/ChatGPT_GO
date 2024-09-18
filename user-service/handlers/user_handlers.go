package handlers

import (
	"ChatGPT_GO/user-service/database"
	"ChatGPT_GO/user-service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// RegisterUser registerUser регистрирует нового пользователя
func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Printf("c.ShouldBindJSON(&newUser): %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.Db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	sqlInsert := `insert into users (username, password) values ($1, $2) returning id`

	err := database.Db.QueryRow(sqlInsert, newUser.Username, newUser.Password).Scan(&newUser.Id)
	if err != nil {
		log.Printf("database.Db.QueryRow(sqlInsert, newUser.Username, newUser.Password).Scan(&newUser.Id) %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
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
