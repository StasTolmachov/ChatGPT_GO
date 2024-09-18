package handlers

import (
	"ChatGPT_GO/task-service/database"
	"ChatGPT_GO/task-service/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// getAllTasks возвращает все задачи в формате JSON
func GetAllTasks(c *gin.Context) {
	rows, err := database.Db.Query("select id, title, description, done from tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
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
func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var task models.Task
	err = database.Db.QueryRow("select id, title, description, done from tasks where id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Done)
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
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlInsert := `insert into tasks (title, description, done) values ($1, $2, $3) returning id`
	err := database.Db.QueryRow(sqlInsert, newTask.Title, newTask.Description, newTask.Done).Scan(&newTask.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

// updateTask обновляет задачу по ID
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlUpdate := `update tasks set title = $1, description = $2, done = $3 where id = $4`

	_, err = database.Db.Exec(sqlUpdate, updateTask.Title, updateTask.Description, updateTask.Done, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateTask)
}

// deleteTask удаляет задачу по ID
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	sqlDelete := `DELETE FROM tasks WHERE id = $1`
	_, err = database.Db.Exec(sqlDelete, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
