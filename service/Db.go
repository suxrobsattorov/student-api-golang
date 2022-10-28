package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "GoDbUser"
	password = "root"
	dbName   = "Go-test-db"
)

func db() *sql.Rows {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	rows, e := db.Query("SELECT * FROM student")
	CheckError(e)
	return rows

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
