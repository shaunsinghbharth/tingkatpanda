package myconnector

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
To create the MySQL DB.
1.user1 (remember to give access to GOLIVEDB to user1)
2.Copy the MySQL Statement and paste it to the Query Window at the MySQL Workbench.

 MySQL Statement
 ---------------
 CREATE TABLE `Items` (
  `ItemId` int NOT NULL AUTO_INCREMENT,
  `ItemName` varchar(256) DEFAULT NULL,
  `ItemPrice` float DEFAULT NULL,
  `ItemDesc` varchar(256) DEFAULT NULL,
  `ItemImg` varchar(256) DEFAULT NULL,
  `ShopID` int DEFAULT NULL,
  PRIMARY KEY (`ItemId`),
  KEY `ShopID_idx` (`ShopID`),
  CONSTRAINT `ShopID` FOREIGN KEY (`ShopID`) REFERENCES `Shops` (`ShopID`) ON DELETE SET NULL ON UPDATE CASCADE
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

func GetItemRecords(db *sql.DB){

	results, err := db.Query("Select * FROM GOLIVEDB.Items")

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

	results, err := db.Exec("INSERT INTO GOLIVEDB.Items(ItemName,ItemPrice,ItemDesc,ItemImg,ShopID) VALUES (?,?,?,?,?)", IN, IP, DE, IG, SI)

	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

//Set Shop ID to Cascade
func EditItemRecord(db *sql.DB, ID int, IN string, IP float64, DE string, IG string, SI int) {

	results, err := db.Exec("UPDATE Items SET ItemName=?, ItemPrice=?, ItemDesc=?, ItemImg=?, ShopID=? WHERE ItemId=?", IN, IP, DE, IG, SI, ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

//Set Shop ID to Null
func DeleteItemRecord(db *sql.DB, ID int) {

	results, err := db.Exec("DELETE FROM Items WHERE ItemId=?", ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func GetSpecificItemRecord(db *sql.DB, ID int) Item{

	results, err := db.Query("Select * FROM GOLIVEDB.Items WHERE ItemId=?", ID)
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
func ConnectItems() sql.DB {
	//dbItems, err := sql.Open("mysql", "tuckmeng:G0L1V3@tcp(128.199.125.231:3306)/GOLIVEDB")
	db, err := sql.Open("mysql", "user1:password@tcp(127.0.0.1:3306)/GOLIVEDB")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened.")
	}
	defer dbItems.Close()

	return *dbItems

}
*/
