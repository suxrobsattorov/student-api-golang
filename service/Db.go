package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "ec2-34-247-72-29.eu-west-1.compute.amazonaws.com"
	port     = 5432
	user     = "sxjmccanruzojc"
	password = "6a47278475f524aec7cb82881271dc283eb1d5476a09d1ce16bafc1bb0d75231"
	dbName   = "decs17qdklniph"
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
