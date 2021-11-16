package dbconfig

import (
	"database/sql"
	"fmt"

	env "crud/utils"

	_ "github.com/go-sql-driver/mysql"
)

var user string = env.GetEnvironmentVars("DBUSER")
var pass string = env.GetEnvironmentVars("DBPASS")
var host string = env.GetEnvironmentVars("DBHOST")
var port string = env.GetEnvironmentVars("DBPORT")
var database string = env.GetEnvironmentVars("DBNAME")

// DB
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func Connect() (*DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, database))
	if err != nil {
		panic(err)
	}
	dbConn.SQL = db
	return dbConn, err
}
