package main 
import (
	"fmt"
	 "log"
	 "database/sql"
	 _ "github.com/mattn/go-sqlite3"
)

func updatePassword(user_name string , new__password string) {
	//Updating the password of the user 
	db,error := sql.Open("sqlite3",database)
	if error != nil {
		log.Fatal(error)
	}
	const password_string = `
	ALTER TABLE USERS SET password = ? WHERE user_name = ?
	`
	_ , new_err := db.Exec(password_string,new__password,user_name)
	if new_err!= nil {
		log.Fatal(error)
	}
	fmt.Println("Updated the password successfully")
	defer db.Close()
}