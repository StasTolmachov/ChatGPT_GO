package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_jwt_secret_key") // Секретный ключ для подписи токенов

// Структура для хранения информации в токене
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Генерация JWT токена для пользователя
func generateToken(username string) (string, error) {
	// Определяем время жизни токена (например, 1 час)
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

// Аутентификация пользователя (логин)
func login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Пример простой проверки имени пользователя и пароля (без базы данных)
	if user.Username == "admin" && user.Password == "password" {
		// Генерация JWT токена
		token, err := generateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Возвращаем JWT токен клиенту
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// Middleware для проверки JWT токена
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		claims := &Claims{}
		// Парсинг и проверка JWT токена
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Сохраняем информацию о пользователе в контексте
		c.Set("username", claims.Username)
		c.Next()
	}
}

// Защищенный ресурс (требует наличия валидного JWT токена)
func protectedResource(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s! You have access to the protected resource.", username)})
}

func main() {
	router := gin.Default()

	// Маршрут для аутентификации (логин)
	router.POST("/login", login)

	// Защищенный маршрут (требует токен)
	router.GET("/protected", authMiddleware(), protectedResource)

	// Запуск сервера
	router.Run(":8080")
}
