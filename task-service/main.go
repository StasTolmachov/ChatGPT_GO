package main

import (
	"ChatGPT_GO/task-service/database"
	"ChatGPT_GO/task-service/handlers"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//todo logger init
	var err error
	database.Db, err = sql.Open("postgres", database.Connection)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Db.Close()

	router := gin.Default()

	router.GET("/tasks/", handlers.GetAllTasks)
	router.GET("/tasks/:id", handlers.GetTaskByID)
	router.POST("/tasks/", handlers.CreateTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)

	router.Run(":8082")
}
