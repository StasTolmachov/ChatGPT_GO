package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

const connStr = "user=postgres dbname=testdb sslmode=disable"

func main() {
	db := ConnectDb(connStr)
	defer db.Close()

	rows := ReadAll(db)
	defer rows.Close()
	ShowAll(rows)

	for {

		fmt.Println("You have options: create, delete, change, done, exit, show, showdone")
		var command string
		fmt.Scan(&command)

		switch command {
		case "exit":
			return
		case "create":
			CreateTask(db)
		case "delete":
			DeleteTask(db)
		case "done":
			DoneTask(db)
		case "show":
			rows := ReadAll(db)
			defer rows.Close()
			ShowAll(rows)
		case "showdone":
			rows := GetDone(db)
			ShowAll(rows)
		default:
			fmt.Println("Unknown command.")

		}

	}

}

// ConnectDb connect to db
func ConnectDb(c string) *sql.DB {

	db, err := sql.Open("postgres", c)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping() // Проверка подключения
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// ReadAll reading all value from db
func ReadAll(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT id, title, description, done FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

// ShowAll show all tasks
func ShowAll(rows *sql.Rows) {
	defer rows.Close()
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task: id: %v Title: %v Description: %v Done: %v\n", task.ID, task.Title, task.Description, task.Done)
	}
}

// create task
func CreateTask(db *sql.DB) {
	sqlInsert := `
   INSERT INTO tasks (title, description)
   VALUES ($1, $2) RETURNING id`

	var input Task
	fmt.Println("write task name:")
	fmt.Scan(&input.Title)
	fmt.Println("write task description:")
	fmt.Scan(&input.Description)

	var id int
	err := db.QueryRow(sqlInsert, input.Title, input.Description).Scan(&id)
	if err != nil {
		fmt.Printf("create task error: %v\n", err)
		return
	}
	fmt.Printf("task create with\nID=%v\ntitle=%v\ndescription=%v\n", id, input.Title, input.Description)
}

// DeleteTask delete task
func DeleteTask(db *sql.DB) {
	sqlDelete := `
DELETE FROM tasks WHERE id = ($1)`
	var id int
	fmt.Println("What id do you want to delete?")
	fmt.Scan(&id)
	_, err := db.Exec(sqlDelete, id)
	if err != nil {
		fmt.Printf("delete task error: %v\n", err)
		return
	}
	fmt.Println("deleted task with id: ", id)

	rows := ReadAll(db)
	defer rows.Close()
	ShowAll(rows)
}

// DoneTask mark done
func DoneTask(db *sql.DB) {
	sqlDone := `
UPDATE tasks SET done = true WHERE id = ($1)`
	fmt.Println("What id do you want to mark done?")
	var id int
	fmt.Scan(&id)
	_, err := db.Exec(sqlDone, id)
	if err != nil {
		fmt.Printf("mark done task error: %v\n", err)
		return
	}
}

// GetDone show all task with status done
func GetDone(db *sql.DB) *sql.Rows {
	sqlShowDone := `
SELECT * FROM tasks WHERE done = true`

	rows, err := db.Query(sqlShowDone)
	if err != nil {
		fmt.Printf("mark done task error: %v\n", err)
	}
	return rows
}
