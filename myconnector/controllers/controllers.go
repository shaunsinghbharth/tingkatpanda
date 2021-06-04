package controllers

import (
	"database/sql"
	"fmt"
	"tingkatpanda/myconnector/models"

	"golang.org/x/crypto/bcrypt"
)

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM GOLIVEDB.Shops")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var shop models.Shop
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

func GetRecordsI(db *sql.DB) {
	results, err := db.Query("Select * FROM GOLIVEDB.Items")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var item models.Item
		err = results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|ItemId:", item.ItemId, "|ItemName:", item.ItemName, "|ItemPrice:", item.ItemPrice, "|ItemDesc:", item.ItemDesc, "|ItemImg:", item.ItemImg, "|ShopId:", item.ShopId, "|")
	}
}

func InsertRecordI(db *sql.DB, ID int, IN string, IP float64, DE string, IG string, SI int) {
	results, err := db.Exec("INSERT INTO GOLIVEDB.Items VALUES (?,?,?,?,?,?)", ID, IN, IP, DE, IG, SI)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func EditRecordI(db *sql.DB, ID int, IN string, IP float64, DE string, IG string, SI int) {
	results, err := db.Exec("UPDATE Shops SET ItemId=?, ItemName=?, ItemPrice=?, ItemDesc=?, ItemDesc=?, ItemImg=?, ShopID=? WHERE ItemId=?", IN, IP, DE, IG, SI, ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func DeleteRecordI(db *sql.DB, ID int) {
	results, err := db.Exec("DELETE FROM Items WHERE ItemId=?", ID)
	if err != nil {
		panic(err)
	} else {
		rows, _ := results.RowsAffected()
		fmt.Println(rows)
	}
}

func GetSpecificRecordI(db *sql.DB, ID int) {
	results, err := db.Query("Select * FROM GOLIVEDB.Items WHERE ItemId=?", ID)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var item models.Item
		err = results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|ItemId:", item.ItemId, "|ItemName:", item.ItemName, "|ItemPrice:", item.ItemPrice, "|ItemDesc:", item.ItemDesc, "|ItemImg:", item.ItemImg, "|ShopId:", item.ShopId, "|")
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
		var shop models.Shop
		err = results.Scan(&shop.ShopId, &shop.ShopName, &shop.ShopAddress, &shop.ShopRating, &shop.ShopPeriod)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("|ShopId:", shop.ShopId, "|ShopName:", shop.ShopName, "|ShopAddress:", shop.ShopAddress, "|ShopRating:", shop.ShopRating, "|ShopPeriod:", shop.ShopPeriod, "|")
	}
}

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
			var person models.User
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

	if err == nil {
		return err == nil
	}
	return err == nil
}
