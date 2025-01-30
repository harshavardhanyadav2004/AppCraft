package signup 
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
const new_Sign_up string = "signup.db"

func createTablesignup(){
	//CREATE TABLE SIGNUP 
	db , err := sql.Open("sqlite3",new_Sign_up)
	if err != nil {
		log.Fatal(err)
	}
	const createtable   string = `
	CREATE TABLE SIGNUP (
	   id INT ,
	   firstname VARCHAR(255),
	   lastname VARCHAR(255),
	   email VARCHAR(255),
	   dob DATE,
	   option VARCHAR(255)
	)
	`
	_ , err = db.Exec(createtable)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created signup table successfully")
	defer db.Close()
}