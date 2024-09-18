package main

import (
	"ChatGPT_GO/user-service/database"
	"ChatGPT_GO/user-service/handlers"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var err error
	database.Db, err = sql.Open("postgres", database.ConnStr)
	if err != nil {
		log.Printf("database.Db, err = sql.Open(\"postgres\", database.ConnStr) %v", err)
		log.Fatal(err)
	}
	if err = database.Db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	defer database.Db.Close()

	router := gin.Default()

	router.POST("/users/register", handlers.RegisterUser)
	router.GET("/users/:id", handlers.GetUserById)

	router.Run(":8082")
}
