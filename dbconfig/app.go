package dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var user string = "root"
var pass string = "1234"
var host string = "127.0.0.1"
var port string = "3306"
var database = "testdb"

// DB ...
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func Connect() (*DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database))
	if err != nil {
		panic(err)
	}
	dbConn.SQL = db
	return dbConn, err
}
