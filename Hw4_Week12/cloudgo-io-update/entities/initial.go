package entities

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var mydb *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:houxi5201314@/todolist_accounts?charset=utf8")
	if err != nil {
		panic(err)
	}
	mydb = db
}

// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// DaoSource Data Access Object Source
type DaoSource struct {
	// if DB, each statement execute sql with random conn.
	// if Tx, all statements use the same conn as the Tx's connection
	SQLExecer
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
