package main

import (
	_ "database/sql"
	"github.com/gofiber/fiber/v2"
)

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
}

func main() {

	app := fiber.New()

	var students []Student

	rows := db()
	defer rows.Close()

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Phone); err != nil {
			//return nil, fmt.Errorf("albumsByArtist : %v", err)
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		//return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	app.Get("/students", func(ctx *fiber.Ctx) error {
		return ctx.JSONP(students)
	})

	//log.Fatal(app.Listen(":8081"))
}
