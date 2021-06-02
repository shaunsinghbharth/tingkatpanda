// Name : Siew Tuck Meng
// Email :tuckmengsiew@gmail.com
package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

/*
To create the MySQL DB.
1.user1 (remember to give access to GOLIVEDB to user1)
2.Copy the MySQL Statement and paste it to the Query Window at the MySQL Workbench.

 MySQL Statement
 ---------------
 CREATE database GOLIVEDB;
 USE GOLIVEDB;
 CREATE TABLE Users (UserName VARCHAR(45) NOT NULL PRIMARY KEY, Password VARCHAR(256));
----------------------------------------------------------------------------------------
3.Run the MySQL Statement by clicking the Thunder Button at the MySQL Workbench.
*/

type User struct {
	UserName string
	Password string
}

//
func insertRecord(db *sql.DB, username, password string) int {
	results, err := db.Exec("INSERT INTO GOLIVEDB.Users VALUES (?,?)", username, password)
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		rows, _ := results.RowsAffected()
		return int(rows)
	}
}

// get the hashed password of the user in string type
func getPasswordOfUser(db *sql.DB, username string) string {
	results, err := db.Query("SELECT * FROM GOLIVEDB.Users where Username=?", username)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		for results.Next() {
			var person User
			err = results.Scan(&person.UserName, &person.Password)
			if err != nil {
				fmt.Println(err)
				return ""
			} else {
				return person.Password
			}
		}
	}
	return ""
}

// hash the given password using bcrypt()
func hashPassword(password string) []byte {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost); err != nil {
		fmt.Println(err)
		return nil
	} else {
		return hash
	}
}

// saved in the db user supplied
func verifyPassword(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return false
	}
	return true
}

func main() {
	db, err := sql.Open("mysql", "user1:password@tcp(127.0.0.1:3306)/GOLIVEDB")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened.")
	}
	defer db.Close()

	for {
		menu := []string{
			"Register New User",
			"Login",
		}

		//switch
		var login int

		//Register New User
		var userNameNew string
		var userPasswordNew string

		//Login
		var userName string
		var userPassword string

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("You have entered an invalid choice. Value should be between %d and %d", 1, len(menu))
			}

		}()

		fmt.Println("User Login Menu: ")
		fmt.Println("================ ")
		for i, g := range menu {
			fmt.Printf("\n%d : %s", i+1, g)
		}
		fmt.Println()
		fmt.Println("What is your choice? ")
		fmt.Scanln(&login)

		fmt.Println("Your Choice is:", menu[login-1])

		switch login {

		case 1:
			//Write a go program that asks the user to enter a username and password.
			//Save the username (trim) and password (hash it) into a DB (show me how to create
			//the DB, show me how to insert a record into the DB using go).

			fmt.Println("1. Register New User")
			fmt.Println("What is the name of the New User?")
			fmt.Scanln(&userNameNew)

			fmt.Println("What is the password of the New User?")
			fmt.Scanln(&userPasswordNew)

			//trim the username
			name := strings.TrimSpace(userNameNew)
			password := userPasswordNew

			//insert the (trim) username and (hash) password into the DB
			fmt.Println(insertRecord(db, name, string(hashPassword(password))))

		case 2:
			//Ask the user to enter a username and password.
			//Validate the username and password to see if the account already exists
			//in the DB.

			fmt.Println("1. Login")
			fmt.Println("What is the User Name?")
			fmt.Scanln(&userName)

			fmt.Println("What is the Password?")
			fmt.Scanln(&userPassword)

			// //---authenticating user---
			//trim the username
			name := strings.TrimSpace(userName)
			password := userPassword

			// retrieve the user's saved password (in string); hashed
			userSavedPassword := getPasswordOfUser(db, name)

			// the password saved in the db the user's supplied password
			if verifyPassword([]byte(userSavedPassword), password) {
				fmt.Println("User authenticated!")
			} else {
				fmt.Println("Invalid username and/or password")
			}

		default:
			fmt.Println("Exit Program")
			break
		}
	}

}
