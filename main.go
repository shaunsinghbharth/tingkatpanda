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

func main() {
	db := myconnector.ConnectShops()

	for {
		menu := []string{
			"Register New User",
			"User Login",
			"Get All User Details",
			"Get Specific User Details",
			"Update An User",
			"Delete An User",
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

			fmt.Println("2. User Login")
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
			//View All Users
			fmt.Println("3.Get All Users Info")

			myconnector.GetApiRecords(&db)

		case 4:
			//Get Specific Users Info
			fmt.Println("4.Get Specific Users Info")

			fmt.Println("What is the specific User Name you want to view?")
			var userName string
			fmt.Scanln(&userName)

			myconnector.GetSpecificApiRecord(&db, userName)

		case 5:
			//Update User
			fmt.Println("5. Update User Details")

			fmt.Println("What is the User Name To Be Updated?")
			fmt.Scanln(&userName)

			fmt.Println("What is the Password To Be Updated?")
			fmt.Scanln(&userPassword)

			myconnector.EditApiRecord(&db, userName, userPassword)

			fmt.Println(userName, userPassword, "are updated.")

		case 6:
			//Delete An User
			fmt.Println("6. Delete An User")

			fmt.Println("What is the User Name to be deleted?")
			fmt.Scanln(&userName)

			myconnector.DeleteApiRecord(&db, userName)

			fmt.Println(userName, "is deleted.")

		case 7:
			//View All Shops
			fmt.Println("7.Get All Shops Info")

			myconnector.GetRecords(&db)

		case 8:
			//Get Specific Shop Info
			fmt.Println("8.Get Specific Shop Info")

			fmt.Println("What is the specific Shop ID you want to view?")
			fmt.Scanln(&idnew)

			myconnector.GetSpecificRecord(&db, idnew)

		case 9:
			//Add New Shop
			fmt.Println("9.Add New Shop")

			fmt.Println("What is the new Shop Name?")
			fmt.Scanln(&namenew)

			fmt.Println("What is the new Shop Address?")
			fmt.Scanln(&addressnew)

			fmt.Println("What is the shop rating of the new shop?")
			fmt.Scanln(&ratingnew)

			fmt.Println("What is the shop availability period of the new shop?")
			fmt.Scanln(&periodnew)

			myconnector.InsertRecordS(&db, namenew, addressnew, ratingnew, periodnew)

			fmt.Println(namenew, addressnew, ratingnew, periodnew, "are added.")

		case 10:
			//Update Shop
			fmt.Println("10. Update Shop Details")

			fmt.Println("What is the Shop ID To Be Updated?")
			fmt.Scanln(&idnew)

			fmt.Println("What is the Shop Name To Be Updated?")
			fmt.Scanln(&namenew)

			fmt.Println("What is the Shop Address To Be Updated?")
			fmt.Scanln(&addressnew)

			fmt.Println("What is the Shop Rating To Be Updated?")
			fmt.Scanln(&ratingnew)

			fmt.Println("What is the Shop Availability Period To Be Updated?")
			fmt.Scanln(&periodnew)

			myconnector.EditRecord(&db, idnew, namenew, addressnew, ratingnew, periodnew)

			fmt.Println(idnew, namenew, addressnew, ratingnew, periodnew, "are updated.")

		case 11:
			//Delete A Shop
			fmt.Println("11. Delete A Shop")

			fmt.Println("What is the Shop ID to be deleted?")
			fmt.Scanln(&idnew)

			myconnector.DeleteRecord(&db, idnew)

			fmt.Println(idnew, "is deleted.")

		case 12:
			//View All Items
			fmt.Println("12.Get All Items Info")

			myconnector.GetItemRecords(&db)

		case 13:
			//Get Specific Item Info
			fmt.Println("13.Get Specific Item Info")

			fmt.Println("What is the specific Item ID you want to view?")
			fmt.Scanln(&iditem)

			myconnector.GetSpecificItemRecord(&db, iditem)

		case 14:
			//Add New Item
			fmt.Println("14.Add New Item")

			fmt.Println("What is the new Item Name?")
			fmt.Scanln(&nameitem)

			fmt.Println("What is the new Item Price?")
			fmt.Scanln(&priceitem)

			fmt.Println("What is the new Item Description?")
			fmt.Scanln(&descitem)

			fmt.Println("What is the new Item Image?")
			fmt.Scanln(&imgitem)

			fmt.Println("What is the new Shop ID?")
			fmt.Scanln(&idshop)

			myconnector.InsertItemRecord(&db, nameitem, priceitem, descitem, imgitem, idshop)

			fmt.Println(nameitem, priceitem, descitem, imgitem, idshop, "are added.")

		case 15:
			//Update Item
			fmt.Println("15. Update Item Details")

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

			myconnector.EditItemRecord(&db, iditem, nameitem, priceitem, descitem, imgitem, idshop)

			fmt.Println(iditem, nameitem, priceitem, descitem, imgitem, idshop, "are updated.")

		case 16:
			//Delete An Item
			fmt.Println("16. Delete An Item")

			fmt.Println("What is the Item ID to be deleted?")
			fmt.Scanln(&iditem)

			myconnector.DeleteItemRecord(&db, iditem)

			fmt.Println(iditem, "is deleted.")

		default:
			fmt.Println("Exit Program")
			break
		}
	}

}
