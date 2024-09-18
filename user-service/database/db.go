package database

import "database/sql"

const ConnStr = "user=postgres dbname=testdb sslmode=disable"

var Db *sql.DB
