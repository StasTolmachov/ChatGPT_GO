package database

import "database/sql"

const Connection = "user=postgres dbname=testdb sslmode=disable"

var Db *sql.DB
