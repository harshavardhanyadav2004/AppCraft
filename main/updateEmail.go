package main 
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func updateEmail(user_name string ,  new_email string  ){
	//Updating the email 
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	} 
     const updation string = `
	 ALTER TABLE USERS SET email = ? user_name = ?	 
	 `
	_,error := db.Exec(updation,new_email,user_name)
	if error != nil {
		log.Fatal(err)
	}
	fmt.Println("Updation of email was done succesfully")
	defer db.Close()
}