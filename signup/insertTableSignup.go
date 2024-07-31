package signup 
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func insertTableSignup(id string , firstname string , lastname string ,email string ,dob string ,option string){
	//Inserting the details of signup details into the databasw 
	db, err := sql.Open("sqlite3", new_Sign_up)
	if err != nil {
		 log.Fatal(err)
	}
	const insertion_string = `
	INSERT INTO SIGNUP (id ,firstname ,lastname , email , dob , option) VALUES (?,?,?,?,?,?)
	`
	_, err = db.Exec(insertion_string, id, firstname, lastname, email, dob,option)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted data into the signup table")
	defer db.Close()
}
func insertIntoUsers(user_name string , email string , password string){
	 //Inserting the details of user into the database
	  db , err:= sql.Open("sqlite3","users.db")
	 if err != nil {
		log.Fatal(err)
	 }
	 const insertion_user_string = `
	 INSERT INTO USERS (id , user_name , email , password ) VALUES (?,?,?)
	 `
	 _, err = db.Exec(insertion_user_string, user_name, email, password)
	 if err != nil {
		log.Fatal(err)
		}
		fmt.Println("Inserted into the users table successfully")
		defer db.Close()
}