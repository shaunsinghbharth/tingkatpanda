package myconnector

import (
	"database/sql"
	"fmt"

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
func InsertRecord(db *sql.DB, username, password string) int {
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
func GetPasswordOfUser(db *sql.DB, username string) string {
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
func HashPassword(password string) []byte {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost); err != nil {
		fmt.Println(err)
		return nil
	} else {
		return hash
	}
}

// saved in the db user supplied
func VerifyPassword(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return false
	}
	return true
}

func Connect() sql.DB {
	db, err := sql.Open("mysql", "tuckmeng:G0L1V3@tcp(128.199.125.231:3306)/GOLIVEDB")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened.")
	}
	defer db.Close()

	return *db

}
