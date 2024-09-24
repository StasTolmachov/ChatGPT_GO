package main

import (
	"ChatGPT_GO/user-service/database"
	"ChatGPT_GO/user-service/handlers"
	"ChatGPT_GO/user-service/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	logger.MakeLogger() //logger init
	logger.Log.Info("start api")

	var err error

	database.Db, err = sql.Open("postgres", database.ConnStr)
	if err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("db is opened")

	if err = database.Db.Ping(); err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("ping db is ok")

	defer database.Db.Close()

	userRepo := &database.PostgresStruct{Db: database.Db}

	router := gin.Default()

	// Маршрут для логина
	router.POST("/users/login", handlers.LoginHandler(userRepo))

	router.POST("/users/register", handlers.RegisterUser(userRepo))
	router.GET("/users/:id", handlers.GetUserById(userRepo))

	router.Run(":8081")
}
