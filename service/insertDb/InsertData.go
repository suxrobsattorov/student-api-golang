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

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	delete := `delete from student`
	_, er := db.Exec(delete)
	CheckError(er)

	insert1 := `insert into "student"("id", "name", "surname", "phone") 
    values(1, 'Suxrob', 'Sattorov', '+998886180111')`
	_, e := db.Exec(insert1)
	CheckError(e)

	insert2 := `insert into "student"("id", "name", "surname", "phone") 
    values($1, $2, $3, $4)`
	_, e = db.Exec(insert2, 2, "Anvar", "Karimov", "+998886183434")
	CheckError(e)

	insert3 := `insert into "student"("id", "name", "surname", "phone") 
    values($1, $2, $3, $4)`
	_, e = db.Exec(insert3, 3, "Jafar", "Narzullayev", "+998886186765")
	CheckError(e)

	insert4 := `insert into "student"("id", "name", "surname", "phone") 
    values($1, $2, $3, $4)`
	_, e = db.Exec(insert4, 4, "Sevara", "Karimova", "+998886184545")
	CheckError(e)

	insert5 := `insert into "student"("id", "name", "surname", "phone") 
    values($1, $2, $3, $4)`
	_, e = db.Exec(insert5, 5, "Gulhayo", "Salimova", "+998886186666")
	CheckError(e)

	insert6 := `insert into "student"("id", "name", "surname", "phone") 
    values($1, $2, $3, $4)`
	_, e = db.Exec(insert6, 6, "Sarvar", "Satimov", "+998886186612")
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
