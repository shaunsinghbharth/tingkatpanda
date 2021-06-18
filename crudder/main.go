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
 CREATE TABLE `Item` (
  `ItemId` int NOT NULL AUTO_INCREMENT,
  `ItemName` varchar(256) DEFAULT NULL,
  `ItemPrice` float DEFAULT NULL,
  `ItemDesc` varchar(256) DEFAULT NULL,
  `ItemImg` varchar(256) DEFAULT NULL,
  `ShopID` int DEFAULT NULL,
  PRIMARY KEY (`ItemId`),
  KEY `ShopID_idx` (`ShopID`),
  CONSTRAINT `ShopID` FOREIGN KEY (`ShopID`) REFERENCES `Shop` (`ShopID`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
----------------------------------------------------------------------------------------
3.Run the MySQL Statement by clicking the Thunder Button at the MySQL Workbench.
*/

type Item struct { // map this type to the record in the Shops table
	ItemId    int
	ItemName  string
	ItemPrice float64
	ItemDesc  string
	ItemImg   string
	ShopId    int //Foreign Key
}

func GetItemRecords(db *sql.DB) {

	results, err := db.Query("Select * FROM GOLIVEDB.Item")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var item Item
		err = results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|ItemId:", item.ItemId, "|ItemName:", item.ItemName, "|ItemPrice:", item.ItemPrice, "|ItemDesc:", item.ItemDesc, "|ItemImg:", item.ItemImg, "|ShopId:", item.ShopId, "|")

	}
}

func InsertItemRecord(db *sql.DB, IN string, IP float64, DE string, IG string, SI int) {

	results, err := db.Exec("INSERT INTO GOLIVEDB.Item(ItemName,ItemPrice,ItemDesc,ItemImg,ShopID) VALUES (?,?,?,?,?)", IN, IP, DE, IG, SI)

	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

//Set Shop ID to Cascade
func EditItemRecord(db *sql.DB, ID int, IN string, IP float64, DE string, IG string, SI int) {

	results, err := db.Exec("UPDATE Item SET ItemName=?, ItemPrice=?, ItemDesc=?, ItemImg=?, ShopID=? WHERE ItemId=?", IN, IP, DE, IG, SI, ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

//Set Shop ID to Null
func DeleteItemRecord(db *sql.DB, ID int) {

	results, err := db.Exec("DELETE FROM Item WHERE ItemId=?", ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func GetSpecificItemRecord(db *sql.DB, ID int) Item {

	results, err := db.Query("Select * FROM GOLIVEDB.Item WHERE ItemId=?", ID)
	if err != nil {
		panic(err.Error())
	}

	var item Item
	for results.Next() {
		// map this type to the record in the table
		err = results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|ItemId:", item.ItemId, "|ItemName:", item.ItemName, "|ItemPrice:", item.ItemPrice, "|ItemDesc:", item.ItemDesc, "|ItemImg:", item.ItemImg, "|ShopId:", item.ShopId, "|")
	}

	return item
}

/*
To create the MySQL DB.
1.user1 (remember to give access to GOLIVEDB to user1)
2.Copy the MySQL Statement and paste it to the Query Window at the MySQL Workbench.

 MySQL Statement
 ---------------
 CREATE database GOLIVEDB;
 USE GOLIVEDB;
 CREATE TABLE `Shop` (
  `ShopId` int NOT NULL AUTO_INCREMENT,
  `ShopName` varchar(256) DEFAULT NULL,
  `ShopAddress` varchar(256) DEFAULT NULL,
  `ShopRating` varchar(45) DEFAULT NULL,
  `ShopPeriod` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`ShopId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
----------------------------------------------------------------------------------------
3.Run the MySQL Statement by clicking the Thunder Button at the MySQL Workbench.
*/

type Shop struct { // map this type to the record in the Shops table
	ShopId      int
	ShopName    string
	ShopAddress string
	ShopRating  string
	ShopPeriod  string
}

func GetShopRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM GOLIVEDB.Shop")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var shop Shop
		err = results.Scan(&shop.ShopId, &shop.ShopName, &shop.ShopAddress, &shop.ShopRating, &shop.ShopPeriod)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|ShopId:", shop.ShopId, "|ShopName:", shop.ShopName, "|ShopAddress:", shop.ShopAddress, "|ShopRating:", shop.ShopRating, "|ShopPeriod:", shop.ShopPeriod, "|")
	}
}

func InsertShopRecordS(db *sql.DB, SN string, SA string, SR string, SP string) {
	results, err := db.Exec("INSERT INTO GOLIVEDB.Shop(ShopName,ShopAddress,ShopRating,ShopPeriod) VALUES (?,?,?,?)", SN, SA, SR, SP)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func EditShopRecord(db *sql.DB, ID int, SN string, SA string, SR string, SP string) {
	results, err := db.Exec("UPDATE Shop SET ShopName=?, ShopAddress=?, ShopRating=?, ShopPeriod=? WHERE ShopId=?", SN, SA, SR, SP, ID)
	fmt.Print(results)

	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func DeleteShopRecord(db *sql.DB, ID int) {
	results, err := db.Exec("DELETE FROM Shop WHERE ShopId=?", ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func GetSpecificShopRecord(db *sql.DB, ID int) {
	results, err := db.Query("Select * FROM GOLIVEDB.Shop WHERE ShopId=?", ID)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var shop Shop
		err = results.Scan(&shop.ShopId, &shop.ShopName, &shop.ShopAddress, &shop.ShopRating, &shop.ShopPeriod)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|ShopId:", shop.ShopId, "|ShopName:", shop.ShopName, "|ShopAddress:", shop.ShopAddress, "|ShopRating:", shop.ShopRating, "|ShopPeriod:", shop.ShopPeriod, "|")
	}
}

/*
To create the MySQL DB.
1.user1 (remember to give access to GOLIVEDB to user1)
2.Copy the MySQL Statement and paste it to the Query Window at the MySQL Workbench.

 MySQL Statement
 ---------------
 CREATE database MYSTOREDB;
 USE MYSTOREDB;
 CREATE TABLE Users (UserName VARCHAR(30) NOT NULL PRIMARY KEY, Password VARCHAR(256));
----------------------------------------------------------------------------------------
3.Run the MySQL Statement by clicking the Thunder Button at the MySQL Workbench.
*/

type User struct {
	UserName string
	Password string
}

//insert the user name and password in the Users table
//func InsertRecord(db *sql.DB, username, password string) int {
func InsertUserRecord(db *sql.DB, username, password string) int {
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

func GetUserRecords(db *sql.DB) {

	results, err := db.Query("Select * FROM GOLIVEDB.Users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var user User
		err = results.Scan(&user.UserName, &user.Password)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|UserName:", user.UserName, "|Password:", user.Password, "|")

	}
}

//Edit User Record
func EditUserRecord(db *sql.DB, UN string, PW string) {

	results, err := db.Exec("UPDATE Users SET Password=? WHERE UserName=?", PW, UN)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

//Delete User Record
func DeleteUserRecord(db *sql.DB, UN string) {

	results, err := db.Exec("DELETE FROM Users WHERE UserName=?", UN)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

//Get Specific User Record
func GetSpecificUserRecord(db *sql.DB, UN string) {

	results, err := db.Query("Select * FROM GOLIVEDB.Users WHERE UserName=?", UN)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var user User
		err = results.Scan(&user.UserName, &user.Password)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|UserName:", user.UserName, "|Password:", user.Password, "|")
	}
}

func main() {
	db, err := sql.Open("mysql", "tuckmeng:G0L1V3@tcp(128.199.125.231:3306)/GOLIVEDB")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened.")
	}
	defer db.Close()

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

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("You have entered an invalid choice. Value should be between %d and %d", 1, len(menu))
			}

		}()

		fmt.Println("Tingkat Panda Menu: ")
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
			//fmt.Println(myconnector.InsertUserRecord(&db, name, string(myconnector.HashPassword(password))))
			fmt.Println(InsertUserRecord(db, name, string(HashPassword(password))))

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
			//userSavedPassword := myconnector.GetPasswordOfUser(&db, name)
			userSavedPassword := GetPasswordOfUser(db, name)

			// the password saved in the db the user's supplied password
			//if myconnector.VerifyPassword([]byte(userSavedPassword), password) {
			if VerifyPassword([]byte(userSavedPassword), password) {
				fmt.Println("User authenticated!")
			} else {
				fmt.Println("Invalid username and/or password")
			}

		case 3:
			//View All Users
			fmt.Println("3.Get All Users Info")

			//myconnector.GetUserRecords(&db)
			GetUserRecords(db)

		case 4:
			//Get Specific Users Info
			fmt.Println("4.Get Specific Users Info")

			fmt.Println("What is the specific User Name you want to view?")
			var userName string
			fmt.Scanln(&userName)

			//myconnector.GetSpecificUserRecord(&db, userName)
			GetSpecificUserRecord(db, userName)

		case 5:
			//Update User
			fmt.Println("5. Update User Details")

			fmt.Println("What is the User Name To Be Updated?")
			fmt.Scanln(&userName)

			fmt.Println("What is the Password To Be Updated?")
			fmt.Scanln(&userPassword)

			//Update the (trim) username and (hash) password into the DB
			name := strings.TrimSpace(userName)
			//password := string(myconnector.HashPassword(userPassword))
			password := string(HashPassword(userPassword))

			//myconnector.EditUserRecord(&db, name, password)
			EditUserRecord(db, name, password)

			fmt.Println(name, password, "are updated.")

		case 6:
			//Delete An User
			fmt.Println("6. Delete An User")

			fmt.Println("What is the User Name to be deleted?")
			fmt.Scanln(&userName)

			//myconnector.DeleteUserRecord(&db, userName)
			DeleteUserRecord(db, userName)

			fmt.Println(userName, "is deleted.")

		case 7:
			//View All Shops
			fmt.Println("7.Get All Shops Info")

			//myconnector.GetShopRecords(&db)
			GetShopRecords(db)

		case 8:
			//Get Specific Shop Info
			fmt.Println("8.Get Specific Shop Info")

			fmt.Println("What is the specific Shop ID you want to view?")
			fmt.Scanln(&idnew)

			//myconnector.GetSpecificShopRecord(&db, idnew)
			GetSpecificShopRecord(db, idnew)

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

			//myconnector.InsertRecordS(&db, namenew, addressnew, ratingnew, periodnew)
			InsertShopRecordS(db, namenew, addressnew, ratingnew, periodnew)

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

			//myconnector.EditShopRecord(&db, idnew, namenew, addressnew, ratingnew, periodnew)
			EditShopRecord(db, idnew, namenew, addressnew, ratingnew, periodnew)

			fmt.Println(idnew, namenew, addressnew, ratingnew, periodnew, "are updated.")

		case 11:
			//Delete A Shop
			fmt.Println("11. Delete A Shop")

			fmt.Println("What is the Shop ID to be deleted?")
			fmt.Scanln(&idnew)

			//myconnector.DeleteShopRecord(&db, idnew)
			DeleteShopRecord(db, idnew)

			fmt.Println(idnew, "is deleted.")

		case 12:
			//View All Items
			fmt.Println("12.Get All Items Info")

			//myconnector.GetItemRecords(&db)
			GetItemRecords(db)

		case 13:
			//Get Specific Item Info
			fmt.Println("13.Get Specific Item Info")

			fmt.Println("What is the specific Item ID you want to view?")
			fmt.Scanln(&iditem)

			//myconnector.GetSpecificItemRecord(&db, iditem)
			GetSpecificItemRecord(db, iditem)

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

			//myconnector.InsertItemRecord(&db, nameitem, priceitem, descitem, imgitem, idshop)
			InsertItemRecord(db, nameitem, priceitem, descitem, imgitem, idshop)

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

			//myconnector.EditItemRecord(&db, iditem, nameitem, priceitem, descitem, imgitem, idshop)
			EditItemRecord(db, iditem, nameitem, priceitem, descitem, imgitem, idshop)

			fmt.Println(iditem, nameitem, priceitem, descitem, imgitem, idshop, "are updated.")

		case 16:
			//Delete An Item
			fmt.Println("16. Delete An Item")

			fmt.Println("What is the Item ID to be deleted?")
			fmt.Scanln(&iditem)

			//myconnector.DeleteItemRecord(&db, iditem)
			DeleteItemRecord(db, iditem)

			fmt.Println(iditem, "is deleted.")

		default:
			fmt.Println("Exit Program")
			break
		}
	}
}
