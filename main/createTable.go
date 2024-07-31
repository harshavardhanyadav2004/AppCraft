package main 
import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)
const database string = "users.db"
func createTable() {
	// Create a Table in the Data Base`
	fmt.Println("Creating the database connection")
	db,error := sql.Open("sql", database) 
	if error != nil {
		log.Fatal(error)
	}
	const creation string = `
	CREATE TABLE USERS (
	   id INT ,
	   user_name VARCHAR(20) NOT NULL ,
	   email VARCHAR(30) UNIQUE NOT NULL,
	   password VARCHAR(10) UNIQUE NOT NULL,
	  primary key (id)
	)
	`
	 _,err := db.Exec(creation)
	 if err!= nil {
		log.Fatal(err)
	 }
	 fmt.Println("Created the Database successfully")
	 defer db.Close()
}
