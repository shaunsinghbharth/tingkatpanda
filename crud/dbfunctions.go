package crud

import (
	"database/sql"
	"fmt"
	"strconv"
	"tingkatpanda/goutils"
)

func GetFullItemRecords(db *sql.DB, itemID string) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Items LEFT JOIN GOLIVEDB.Shops ON Items.ShopID = Shops.ShopID WHERE Items.ItemId = ? UNION Select * FROM GOLIVEDB.Items RIGHT JOIN GOLIVEDB.Shops ON Items.ShopID = Shops.ShopID WHERE Items.ItemId = ?", itemID, itemID)

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func GetFullItemRecordsAll(db *sql.DB) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Items LEFT JOIN GOLIVEDB.Shops ON Items.ShopID = Shops.ShopID")

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func GetUserRecords(db *sql.DB) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Users")

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func GetUserItemsRecords(db *sql.DB, userID string) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.UserItems WHERE UserName=?", userID)

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)
	fmt.Println("GetUserItems ", returnMaps)

	return returnMaps
}

func GetShopRecords(db *sql.DB) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Shops")

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func GetItemRecords(db *sql.DB) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Items")

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func GetSpecificUserRecords(db *sql.DB, AN string) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Users WHERE UserName=?", AN)

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func GetSpecificShopRecords(db *sql.DB, SI string) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Shops WHERE ShopID=?", SI)

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func GetSpecificItemRecords(db *sql.DB, II string) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Items WHERE ItemID=?", II)

	if err != nil {
		panic(err.Error())
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func EditUserRecords(db *sql.DB, AN string, AK string) []map[string]interface{} {
	results, err := db.Query("UPDATE GOLIVEDB.Users SET Password=? WHERE UserName=?", AK, AN)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Users WHERE UserName=?", AN)
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

func EditShopRecords(db *sql.DB, shopID string, shopName string, shopAddress string, shopRating string, shopStart string, shopEnd string, shopPostCode string) []map[string]interface{} {
	//func EditShopRecords(db *sql.DB, ID int, SN string, SA string, SR string, SP string) []Shops {
	IDint, _ := strconv.Atoi(shopID)

	fmt.Println("EDITSHOPSDB")
	_ , err := db.Exec("UPDATE GOLIVEDB.Shops SET ShopName=?, ShopAddress=?, ShopRating=?, ShopPostalCode = ? WHERE ShopID=?", shopName, shopAddress, shopRating, shopPostCode, IDint)

	if err != nil {
		panic(err.Error())
	} else {
		//results, err = db.Query("Select * FROM GOLIVEDB.Shops WHERE ShopID=?", shopID)
	}

	//returnMaps := goutils.SQLtoMap(results)

	return nil
}

func CreateShopRecords(db *sql.DB, shopID string, shopName string, shopAddress string, shopRating string, shopStart string, shopEnd string, shopPostCode string) []map[string]interface{} {
	//func EditShopRecords(db *sql.DB, ID int, SN string, SA string, SR string, SP string) []Shops {

	fmt.Println("EDITSHOPSDB")
	_ , err := db.Exec("INSERT INTO GOLIVEDB.Shops (ShopName, ShopAddress, ShopRating, ShopPostalCode) VALUES (?,?,?,?)", shopName, shopAddress, shopRating, shopPostCode)

	if err != nil {
		panic(err.Error())
	} else {
		//results, err = db.Query("Select * FROM GOLIVEDB.Shops WHERE ShopID=?", shopID)
	}

	//returnMaps := goutils.SQLtoMap(results)

	return nil
}

func EditItemRecords(db *sql.DB, ItemID string, ItemName string, ItemCategory, ItemPrice string, ItemDescription string, ItemImage string, ItemTiming string, ShopID string) []map[string]interface{} {

	fmt.Println("VARS ", ItemID, ItemCategory, ItemPrice, ItemTiming, ItemName, ShopID)
	itemIDint, _ := strconv.Atoi(ItemID)
	//shopIDint, _ := strconv.Atoi(ShopID)
	itemTimingint, _ := strconv.Atoi(ItemTiming)

	//tx, err := db.Begin()
	res, err := db.Exec("UPDATE GOLIVEDB.Items SET ItemName=?, ItemCategory=?, ItemPrice=?, ItemDesc=?, ItemImg=?, ItemTiming=? WHERE ItemId=?", ItemName, ItemCategory, ItemPrice, ItemDescription, ItemImage, itemTimingint, itemIDint)
	//tx.Commit()

	rows, err := res.RowsAffected()
	fmt.Println("Rows Affected: ", rows)
	if err != nil {
		panic(err.Error())
	} else {
		//results, err = db.Query("Select * FROM GOLIVEDB.Items WHERE ItemID=?", ItemID)
	}

	return nil
}

func CreateItemRecords(db *sql.DB, ItemID string, ItemName string, ItemCategory, ItemPrice string, ItemDescription string, ItemImage string, ItemTiming string, ShopID string) []map[string]interface{} {

	fmt.Println("VARS ", ItemID, ItemCategory, ItemPrice, ItemTiming, ItemName, ShopID)
	//itemIDint, _ := strconv.Atoi(ItemID)
	shopIDint, _ := strconv.Atoi(ShopID)
	itemTimingint, _ := strconv.Atoi(ItemTiming)

	//tx, err := db.Begin()
	res, err := db.Exec("INSERT INTO GOLIVEDB.Items (ShopId, ItemName, ItemCategory, ItemPrice, ItemDesc, ItemImg, ItemTiming) VALUES (?, ?, ?, ?, ?, ?, ?)", shopIDint, ItemName, ItemCategory, ItemPrice, ItemDescription, ItemImage, itemTimingint)
	//tx.Commit()

	rows, err := res.RowsAffected()
	fmt.Println("Rows Affected: ", rows)
	if err != nil {
		panic(err.Error())
	} else {
		//results, err = db.Query("Select * FROM GOLIVEDB.Items WHERE ItemID=?", ItemID)
	}

	return nil
}

func DeleteUserRecords(db *sql.DB, AN string) string {
	results, err := db.Query("DELETE FROM Users WHERE UserName=?", AN)

	if err != nil {
		return "User Does Not Exist"
		//panic(err.Error())
	}

	if results != nil {
		return "Error Deleting"
	}

	return "Delete Success"
}

func DeleteShops(db *sql.DB, ID string) string {
	IDint, _ := strconv.Atoi(ID)
	fmt.Println("DELETESHOPSDB", IDint)
	results, err := db.Exec("DELETE FROM Shops WHERE ShopId=?", IDint)

	if err != nil {
		return "Shop Does Not Exist"
		//panic(err.Error())
	}

	if results != nil {
		return "Error Deleting"
	}

	return "Delete Success"
}

func DeleteItemRecords(db *sql.DB, ID string) string {
	fmt.Println("DELETEITEMSUSER ", ID)
	ID_int, _ := strconv.Atoi(ID)
	//results, err := db.Query("DELETE FROM Items WHERE ItemID=?", ID_int)
	results, err := db.Exec("DELETE FROM Items WHERE ItemID=?", ID_int)

	if err != nil {
		return "Item Does Not Exist"
		//panic(err.Error())
	}

	rows, err := results.RowsAffected()
	fmt.Println("ROWS AFFECTED: ", rows)

	return "Delete Success"
}
