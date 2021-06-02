// Name : Siew Tuck Meng
// Email :tuckmengsiew@gmail.com
package main

import (
	"fmt"
	"strings"
	"tingkatpanda/myconnector"
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

func main() {
	db := myconnector.Connect()

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
			fmt.Println(myconnector.InsertRecord(&db, name, string(myconnector.HashPassword(password))))

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
			userSavedPassword := myconnector.GetPasswordOfUser(&db, name)

			// the password saved in the db the user's supplied password
			if myconnector.VerifyPassword([]byte(userSavedPassword), password) {
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
