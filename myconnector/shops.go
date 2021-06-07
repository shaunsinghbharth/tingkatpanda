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
 CREATE database GOLIVEDB;
 USE GOLIVEDB;
 CREATE TABLE `Shops` (
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

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM GOLIVEDB.Shops")

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

func InsertRecordS(db *sql.DB, SN string, SA string, SR string, SP string) {
	results, err := db.Exec("INSERT INTO GOLIVEDB.Shops(ShopName,ShopAddress,ShopRating,ShopPeriod) VALUES (?,?,?,?)", SN, SA, SR, SP)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func EditRecord(db *sql.DB, ID int, SN string, SA string, SR string, SP string) {
	results, err := db.Exec("UPDATE Shops SET ShopName=?, ShopAddress=?, ShopRating=?, ShopPeriod=? WHERE ShopId=?", SN, SA, SR, SP, ID)
	fmt.Print(results)

	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func DeleteRecord(db *sql.DB, ID int) {
	results, err := db.Exec("DELETE FROM Shops WHERE ShopId=?", ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func GetSpecificRecord(db *sql.DB, ID int) {
	results, err := db.Query("Select * FROM GOLIVEDB.Shops WHERE ShopId=?", ID)
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

func ConnectShops() sql.DB {
	dbS, err := sql.Open("mysql", "tuckmeng:G0L1V3@tcp(128.199.125.231:3306)/GOLIVEDB")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened.")
	}
	defer dbS.Close()

	return *dbS

}
