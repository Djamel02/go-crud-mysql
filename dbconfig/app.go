package dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var user string = GetEnvironmentVars("DBUSER")
var pass string = GetEnvironmentVars("DBPASS")
var host string = GetEnvironmentVars("DBHOST")
var port string = GetEnvironmentVars("DBPORT")
var database string = GetEnvironmentVars("DBNAME")

// DB
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
