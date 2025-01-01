package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created or already exists.")

	insertUserSQL := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err = db.Exec(insertUserSQL, "Alice", 25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user Alice")

	_, err = db.Exec(insertUserSQL, "Bob", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Bob.")

	rows, err := db.Query(`SELECT id, name, age FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users in the database:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
