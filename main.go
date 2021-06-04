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

type Shop struct { // map this type to the record in the Shops table
	ShopId      int
	ShopName    string
	ShopAddress string
	ShopRating  string
	ShopPeriod  string
}

type Item struct { // map this type to the record in the Shops table
	ItemId    int
	ItemName  string
	ItemPrice float64
	ItemDesc  string
	ItemImg   string
	ShopId    int
}

func main() {
	db := myconnector.Connect()

	for {
		menu := []string{
			"Register New User",
			"Login",
			"Get All Shops Details",
			"Get Specific Shop Details",
			"Add New Shop",
			"Update Shop",
			"Delete A Shop",
			"Get All Items Details",
			"Get Specific Item Details",
			"Add New Item",
			"Update Item",
			"Delete An Item",
		}

		//switch
		var login int

		//Register New User
		var userNameNew string
		var userPasswordNew string

		//Login
		var userName string
		var userPassword string

		//Shops Table
		var idnew int
		var namenew string
		var addressnew string
		var ratingnew string
		var periodnew string

		//Items Table
		var iditem int
		var nameitem string
		var priceitem float64
		var descitem string
		var imgitem string
		var idshop int

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

		case 3:
			//View All Shops
			fmt.Println("1.Get All Shops Info")

			myconnector.GetRecords(db)

		case 4:
			//Get Specific Shop Info
			fmt.Println("2.Get Specific Shop Info")

			fmt.Println("What is the specific Shop you want to view?")
			fmt.Scanln(&idnew)

			myconnector.GetSpecificRecord(db, idnew)

		case 5:
			//Add New Shop
			fmt.Println("2.Add New Shop")

			fmt.Println("Shop ID is auto generated?")
			fmt.Scanln(&idnew)

			fmt.Println("What is the new Shop Address?")
			fmt.Scanln(&namenew)

			fmt.Println("What is the new Shop Address?")
			fmt.Scanln(&addressnew)

			fmt.Println("What is the shop rating of the new shop?")
			fmt.Scanln(&ratingnew)

			fmt.Println("What is the shop availability period of the new shop?")
			fmt.Scanln(&periodnew)

			myconnector.InsertRecord(db, idnew, namenew, addressnew, ratingnew, periodnew)

			fmt.Println(idnew, namenew, addressnew, ratingnew, periodnew, "is added.")

		case 6:
			//Update Shop
			fmt.Println("4. Update Shop Details")

			fmt.Println("What is the Shop To Be Updated?")
			fmt.Scanln(&idnew)

			fmt.Println("What is the Shop Name To Be Updated?")
			fmt.Scanln(&namenew)

			fmt.Println("What is the Shop Address To Be Updated?")
			fmt.Scanln(&addressnew)

			fmt.Println("What is the Shop Rating To Be Updated?")
			fmt.Scanln(&ratingnew)

			fmt.Println("What is the Shop Availability Period To Be Updated?")
			fmt.Scanln(&periodnew)

			myconnector.EditRecord(db, idnew, namenew, addressnew, ratingnew, periodnew)

			fmt.Println(idnew, namenew, addressnew, ratingnew, periodnew, "are updated.")

		case 7:
			//Delete A Shop
			fmt.Println("5. Delete A Shop")

			fmt.Println("What is the Shop to be deleted?")
			fmt.Scanln(&idnew)

			myconnector.DeleteRecord(db, idnew)

			fmt.Println(idnew, "is deleted.")

		case 8:
			//View All Items
			fmt.Println("1.Get All Items Info")

			myconnector.GetRecordsI(db)

		case 9:
			//Get Specific Item Info
			fmt.Println("2.Get Specific Item Info")

			fmt.Println("What is the specific Item you want to view?")
			fmt.Scanln(&iditem)

			myconnector.GetSpecificRecordI(db, iditem)

		case 10:
			//Add New Item
			fmt.Println("2.Add New Item")

			fmt.Println("Item ID is auto generated?")
			fmt.Scanln(&iditem)

			fmt.Println("What is the new Item Name?")
			fmt.Scanln(&nameitem)

			fmt.Println("What is the new Item Price?")
			fmt.Scanln(&priceitem)

			fmt.Println("What is the new Item Description?")
			fmt.Scanln(&descitem)

			fmt.Println("What is the new Item Image?")
			fmt.Scanln(&imgitem)

			fmt.Println("What is the new Shop ID To Be Updated?")
			fmt.Scanln(&idshop)

			myconnector.InsertRecordI(db, iditem, nameitem, priceitem, descitem, imgitem, idshop)

			fmt.Println(idnew, namenew, addressnew, ratingnew, periodnew, "is added.")

		case 11:
			//Update Item
			fmt.Println("4. Update Item Details")

			fmt.Println("What is the Item ID To Be Updated?")
			fmt.Scanln(&iditem)

			fmt.Println("What is the Item Name To Be Updated?")
			fmt.Scanln(&nameitem)

			fmt.Println("What is the Item Price To Be Updated?")
			fmt.Scanln(&priceitem)

			fmt.Println("What is the Item Description To Be Updated?")
			fmt.Scanln(&descitem)

			fmt.Println("What is the Item Image To Be Updated?")
			fmt.Scanln(&imgitem)

			fmt.Println("What is the Shop ID To Be Updated?")
			fmt.Scanln(&idshop)

			myconnector.EditRecordI(db, iditem, nameitem, priceitem, descitem, imgitem, idshop)

			fmt.Println(iditem, nameitem, priceitem, descitem, imgitem, idshop, "are updated.")

		case 12:
			//Delete An Item
			fmt.Println("5. Delete An Item")

			fmt.Println("What is the Item ID to be deleted?")
			fmt.Scanln(&iditem)

			myconnector.DeleteRecordI(db, iditem)

			fmt.Println(iditem, "is deleted.")

		default:
			fmt.Println("Exit Program")
			break
		}
	}

}
