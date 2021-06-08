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
 CREATE TABLE `ApiUsers` (
  `UserName` varchar(45) NOT NULL,
  `ApiKey` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`UserName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

----------------------------------------------------------------------------------------
3.Run the MySQL Statement by clicking the Thunder Button at the MySQL Workbench.
*/

type Api struct { // map this type to the record in the Shops table
	ApiName string
	ApiKeys string
}

func GetApiRecords(db *sql.DB) {

	results, err := db.Query("Select * FROM GOLIVEDB.ApiUsers")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var api Api
		err = results.Scan(&api.ApiName, &api.ApiKeys)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|UserName:", api.ApiName, "|ApiKey:", api.ApiKeys, "|")

	}
}

//Edit Api Key
func EditApiRecord(db *sql.DB, AN string, AK string) {

	results, err := db.Exec("UPDATE ApiUsers SET ApiKey=? WHERE UserName=?", AK, AN)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

//Delete Api Key
func DeleteApiRecord(db *sql.DB, AN string) {

	results, err := db.Exec("DELETE FROM ApiUsers WHERE UserName=?", AN)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func GetSpecificApiRecord(db *sql.DB, AN string) {

	results, err := db.Query("Select * FROM GOLIVEDB.ApiUsers WHERE UserName=?", AN)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var api Api
		err = results.Scan(&api.ApiName, &api.ApiKeys)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|UserName:", api.ApiName, "|ApiKey:", api.ApiKeys, "|")
	}
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
