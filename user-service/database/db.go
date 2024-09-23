package database

import (
	"ChatGPT_GO/user-service/logger"
	"database/sql"
)

const ConnStr = "user=postgres dbname=testdb sslmode=disable host=localhost port=5432"

var Db *sql.DB

type DbInterface interface {
	DbPingMethod() error
	RegisterUserMethod(username, password string) (int, error)
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
	r.Db.QueryRow(sqlInsert, username, password).Scan(&id)
	logger.Log.Info()
	return id, nil
}
