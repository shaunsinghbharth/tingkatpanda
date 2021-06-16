package crud

import (
	"database/sql"
	"fmt"
	"tingkatpanda/goutils"
)

func GetFullItemRecords(db *sql.DB, itemID string) []map[string]interface{} {
	results, err := db.Query("Select * FROM GOLIVEDB.Items LEFT JOIN GOLIVEDB.Shops ON Items.ShopID = Shops.ShopID WHERE Items.ItemId = ? UNION Select * FROM GOLIVEDB.Items RIGHT JOIN GOLIVEDB.Shops ON Items.ShopID = Shops.ShopID WHERE Items.ItemId = ?", itemID, itemID)

	if err != nil {
		panic(err.Error())
	}

	/*
		for results.Next() {
			var item CombinedItem
			results.Scan(&item.Item.ItemId, &item.Item.ItemName, &item.Item.ItemDesc, &item.Item.ItemPrice, &item.Item.ItemImg, &item.Item.ShopId,
				&item.Shop.ShopId, &item.Shop.ShopName, &item.Shop.ShopAddress, &item.Shop.ShopStart, &item.Shop.ShopEnd, &item.Shop.ShopRating)
			returnVal = append(returnVal, item)
		}*/

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

//EditShopRecords(&db, shopID, shopName, shopAddress, shopRating, shopStart, shopEnd, shopPostCode)
func EditShopRecords(db *sql.DB, shopID string, shopName string, shopAddress string, shopRating string, shopStart string, shopEnd string, shopPostCode string) []map[string]interface{} {
	//func EditShopRecords(db *sql.DB, ID int, SN string, SA string, SR string, SP string) []Shops {
	results, err := db.Query("UPDATE Shops SET ShopName=?, ShopAddress=?, ShopRating=?, ShopStart=?, ShopEnd = ?, ShopPostCode = ? WHERE ShopId=?", shopName, shopAddress, shopRating, shopStart, shopEnd, shopPostCode, shopID)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Shops WHERE ShopID=?", shopID)
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
}

//EditItemRecords(&db, itemID, itemName, itemPrice, itemDesc, itemImg, shopID)
func EditItemRecords(db *sql.DB, ItemID string, ItemName string, ItemPrice string, ItemDescription string, ItemImage string, ShopID string) []map[string]interface{} {
	results, err := db.Query("UPDATE Items SET ItemName=?, ItemPrice=?, ItemDesc=?, ItemImg=?, ShopID=? WHERE ItemId=?", ItemName, ItemPrice, ItemDescription, ItemImage, ShopID, ItemID)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Items WHERE ItemID=?", ItemID)
	}

	returnMaps := goutils.SQLtoMap(results)

	return returnMaps
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

func DeleteShopRecords(db *sql.DB, ID string) string {
	results, err := db.Query("DELETE FROM Shops WHERE ShopID=?", ID)

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
	results, err := db.Query("DELETE FROM Items WHERE ItemID=?", ID)

	if err != nil {
		return "Item Does Not Exist"
		//panic(err.Error())
	}

	if results != nil {
		return "Error Deleting"
	}

	return "Delete Success"
}
