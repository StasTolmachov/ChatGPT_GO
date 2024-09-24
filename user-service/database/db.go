package database

import (
	"ChatGPT_GO/user-service/logger"
	"ChatGPT_GO/user-service/models"
	"database/sql"
)

const ConnStr = "user=postgres dbname=testdb sslmode=disable host=localhost port=5432"

var Db *sql.DB

type DbInterface interface {
	DbPingMethod() error
	RegisterUserMethod(username, password string) (int, error)
	GetUserByIdMethod(id int) (string, error)
	LoginHandlerMethod(u string) (int, string, error)
	CheckUserExistMethod(u string) (int, error)
}

type PostgresStruct struct {
	Db *sql.DB
}

func (r *PostgresStruct) DbPingMethod() error {
	logger.Log.Info()
	return r.Db.Ping()
}

func (r *PostgresStruct) RegisterUserMethod(username, password string) (int, error) {
	var id int
	sqlInsert := `insert into users (username, password) values ($1, $2) returning id`
	err := r.Db.QueryRow(sqlInsert, username, password).Scan(&id)
	logger.Log.Info()
	return id, err
}

func (r *PostgresStruct) GetUserByIdMethod(id int) (string, error) {
	var user models.User
	// Выполняем SQL-запрос
	err := r.Db.QueryRow("SELECT id, username FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username)
	logger.Log.Info()
	return user.Username, err
}
func (r *PostgresStruct) LoginHandlerMethod(u string) (int, string, error) {
	var user models.User
	err := r.Db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", u).
		Scan(&user.Id, &user.Password)
	logger.Log.Info()
	return user.Id, user.Password, err
}

func (r *PostgresStruct) CheckUserExistMethod(u string) (int, error) {
	var user models.User
	err := r.Db.QueryRow("SELECT id FROM users WHERE username = $1", u).
		Scan(&user.Id)
	logger.Log.Info()
	return user.Id, err
}
