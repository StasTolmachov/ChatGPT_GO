package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Done        bool   `json:"Done"`
}

const connection = "user=postgres dbname=testdb sslmode=disable"

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/api/tasks", getAllTasks)
	router.GET("/api/tasks/:id", getTaskByID)
	router.POST("/api/tasks", createTask)
	router.PUT("/api/tasks/:id", updateTask)
	router.DELETE("/api/tasks/:id", deleteTask)

	router.Run(":8080")

}

// getAllTasks возвращает все задачи в формате JSON
func getAllTasks(c *gin.Context) {
	rows, err := db.Query("select id, title, description, done from tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}

// getTaskByID возвращает задачу по ID
func getTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var task Task
	err = db.QueryRow("select id, title, description, done from tasks where id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Done)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, task)
}

// createTask создает новую задачу
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlInsert := `insert into tasks (title, description, done) values ($1, $2, $3) returning id`
	err := db.QueryRow(sqlInsert, newTask.Title, newTask.Description, newTask.Done).Scan(&newTask.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

// updateTask обновляет задачу по ID
func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var updateTask Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlUpdate := `update tasks set title = $1, description = $2, done = $3 where id = $4`

	_, err = db.Exec(sqlUpdate, updateTask.Title, updateTask.Description, updateTask.Done, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateTask)
}

// deleteTask удаляет задачу по ID
func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	sqlDelete := `DELETE FROM tasks WHERE id = $1`
	_, err = db.Exec(sqlDelete, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
