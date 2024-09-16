package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

const connStr = "user=postgres dbname=testdb sslmode=disable"

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	router.Use(Logger())

	router.LoadHTMLGlob("templates/*")

	router.GET("/tasks", Home)
	router.POST("/createTask", CreateTask)
	router.POST("/tasks/:id/delete", deleteTask)
	router.POST("/tasks/:id/done", markDone)

	router.Run(":8080")

}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Обрабатываем запрос
		c.Next()

		// Логируем время выполнения
		latency := time.Since(t)
		log.Printf("Request took %v", latency)
	}
}

func Home(c *gin.Context) {
	tasks := GetAllTasks()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Tasks": tasks,
	})
}

// CreateTask create task
func CreateTask(c *gin.Context) {

	Title := c.PostForm("Title")
	Description := c.PostForm("Description")

	err := CreateTaskDB(Title, Description)
	if err != nil {
		log.Printf("CreateTask error: %v\n", err)
	}

	c.Redirect(http.StatusFound, "/tasks")

}

// deleteTask create task
func deleteTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("deleteTask error: %v\n", err)
		return
	}
	err = DeleteTaskDb(id)
	if err != nil {
		log.Printf("DeleteTaskDb error: %v\n", err)
	}

	c.Redirect(http.StatusFound, "/tasks")
}

// markDone create task
func markDone(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("markDone: %v\n", err)
		return
	}
	err = markDoneDb(id)

	c.Redirect(http.StatusFound, "/tasks")

}

// markDoneDb mark done
func markDoneDb(id int) error {
	sqlDone := `
UPDATE tasks SET done = true WHERE id = ($1)`

	_, err := db.Exec(sqlDone, id)
	return err
}

// DeleteTaskDb delete task
func DeleteTaskDb(id int) error {
	sqlDelete := `
DELETE FROM tasks WHERE id = ($1)`

	_, err := db.Exec(sqlDelete, id)
	return err

}

// CreateTaskDB CreateTask create task
func CreateTaskDB(t, d string) error {
	sqlInsert := `
   INSERT INTO tasks (title, description)
   VALUES ($1, $2) RETURNING id`

	err := db.QueryRow(sqlInsert, t, d).Err()
	return err

}

// GetAllTasks get all value from db and return
func GetAllTasks() []Task {
	rows, err := db.Query("SELECT id, title, description, done FROM tasks")
	if err != nil {
		log.Fatal("GetAllTasks error: ", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done)
		if err != nil {
			log.Fatal("Error scanning task: ", err)
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error during row iteration: ", err)
	}
	return tasks
}
