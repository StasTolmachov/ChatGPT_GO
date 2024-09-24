package handlers

import (
	"ChatGPT_GO/user-service/auth"
	"ChatGPT_GO/user-service/database"
	"ChatGPT_GO/user-service/logger"
	"ChatGPT_GO/user-service/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RegisterUser registerUser регистрирует нового пользователя
func RegisterUser(repo database.DbInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		var err error

		// Привязка данных JSON
		if err := c.ShouldBindJSON(&newUser); err != nil {
			logger.Log.Errorf("Привязка данных JSON %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.DbPingMethod(); err != nil {
			logger.Log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
			return
		}
		// проверяем есть имя пользователя в базе
		newUser.Id, err = repo.CheckUserExistMethod(newUser.Username)
		logger.Log.Info(newUser)

		//если есть ошибка значит юзера нет
		if err != nil {
			if err == sql.ErrNoRows { //проверяем соответствие ошибки что нет нужной строки
				//создаем нового юзера
				id, err := repo.RegisterUserMethod(newUser.Username, newUser.Password)
				if err != nil {
					logger.Log.Errorf("значит юзера нет %s", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				newUser.Id = id
				c.JSON(http.StatusCreated, newUser)
				logger.Log.Infof("юзера нет потому создали: id=%v user=%v pass=%v", newUser.Id, newUser.Username, newUser.Password)
			} else { //если ошибка другого рода выводим ее
				logger.Log.Errorf("неизвестная ошибка при проверке наличия юзера: %s", err)
			}

		} else { //если нет ошибки значит юзер есть и выводим логин
			c.JSON(http.StatusFound, newUser)
			logger.Log.Infof("User exist: id=%v user=%v pass=%v", newUser.Id, newUser.Username, newUser.Password)
		}

	}

}

func GetUserById(repo database.DbInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var err error
		user.Id, err = strconv.Atoi(c.Param("id"))

		if database.Db == nil {
			logger.Log.Fatal("database.Db == nil")
		}

		user.Username, err = repo.GetUserByIdMethod(user.Id)
		if err != nil {
			logger.Log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка выполнения запроса"})
			return
		}

		c.JSON(http.StatusOK, user)
		logger.Log.Info(user)
	}

}

// LoginHandler — обработчик логина
func LoginHandler(repo database.DbInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Привязка данных JSON
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			logger.Log.Errorf("Привязка данных JSON %s", err)
			return
		}

		// Проверка пользователя в базе данных
		var user models.User
		var err error
		user.Id, user.Password, err = repo.LoginHandlerMethod(loginData.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
				return
			}
			logger.Log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}

		// Валидация пароля
		if loginData.Password != user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			logger.Log.Error("Invalid username or password")
			return
		}

		// Генерация JWT токена
		token, err := auth.GenerateJWT(user.Username)
		if err != nil {
			logger.Log.Errorf("Генерация JWT токена %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		// Возвращаем токен клиенту
		c.JSON(http.StatusOK, gin.H{"token": token})
	}

}

func ProtectedRoute(c *gin.Context) {
	// Извлекаем имя пользователя из токена
	username := c.MustGet("username").(string)

	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + username + "!"})
}
