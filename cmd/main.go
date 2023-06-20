package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}
	fmt.Println("sera")
	db()
}

func db() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/app_development")
	if err != nil {
		// handle the error
		log.Print(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// handle the error
	}
	// rows, err := db.Query("SELECT * FROM Usuario")
	// if err != nil {
	// 	// handle the error
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	// scan the values from the row into variables
	// 	err = rows.Scan(&id, &name)
	// 	if err != nil {
	// 		// handle the error
	// 	}
	// 	// use the retrieved values
	// 	fmt.Println(id, name)
	// }

	// if err = rows.Err(); err != nil {
	// 	// handle the error
	// }
}
