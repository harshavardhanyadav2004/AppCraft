package main 
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func  updateUserName(user_name string , new_user_name string)  {
	 //Update the User Name 
	 db, err := sql.Open("sqlite3", database)
	 if err!=nil {
		log.Fatal(err)
	 }
	 const updation_string = `
	 ALTER TABLE USERS SET user_name = ? WHERE user_name = ?
	 `
	 _ , error := db.Exec(updation_string,user_name,new_user_name)
	 if error != nil {
		 log.Fatal(err)
	 }
	 fmt.Println("Updation was done successfully")
	 defer db.Close()
}
