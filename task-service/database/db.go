package database

import "database/sql"

const Connection = "user=postgres dbname=testdb password=1234 sslmode=disable host=postgres port=5432"

var Db *sql.DB
