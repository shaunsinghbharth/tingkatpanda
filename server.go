package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tingkatpanda/goutils"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct { // map this type to the record in the Users table
	UserName string `json:"username"`
	Password string `json:"password"`
}

var db sql.DB

func init() {
	db = Connect()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":5555", router))
}

func initaliseHandlers(router *mux.Router) {
	router.Use(Middleware)

	router.HandleFunc("/getfullitem", GetFullItem).Methods("GET")

	router.HandleFunc("/getusers", GetAllUser).Methods("GET")
	router.HandleFunc("/getspecificusers", GetUserByUserName).Methods("GET")
	router.HandleFunc("/editusers", EditUser).Methods("GET")
	router.HandleFunc("/deleteusers", DeleteUser).Methods("GET")

	router.HandleFunc("/getshops", GetAllShop).Methods("GET")
	router.HandleFunc("/getspecificshops", GetShopByShopID).Methods("GET")
	router.HandleFunc("/editshops", EditShop).Methods("GET")
	router.HandleFunc("/deleteshops", DeleteShop).Methods("GET")

	router.HandleFunc("/getitems", GetAllItem).Methods("GET")
	router.HandleFunc("/getspecificitems", GetItemByItemID).Methods("GET")
	router.HandleFunc("/edititems", EditItem).Methods("GET")
	router.HandleFunc("/deleteitems", DeleteItem).Methods("GET")

}

//GetAllUser get all user data
func GetFullItem(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("item")

	var item []map[string]interface{}
	item = GetFullItemRecords(&db, param1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//GetAllUser get all user data
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	var user []map[string]interface{}
	user = GetUserRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
}

//GetAllShop get all shop data
func GetAllShop(w http.ResponseWriter, r *http.Request) {

	var shop []map[string]interface{}
	shop = GetShopRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)

	fmt.Println("SUCCESS")
}

//GetAllItem get all item data
func GetAllItem(w http.ResponseWriter, r *http.Request) {

	var item []map[string]interface{}
	item = GetItemRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//GetUserByUserName returns user with specific UserName
func GetUserByUserName(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")

	var user []map[string]interface{}
	user = GetSpecificUserRecords(&db, userID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
}

//GetShopByShopID returns shop with specific ShopID
func GetShopByShopID(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Query().Get("shopid")

	var shop []map[string]interface{}
	shop = GetSpecificShopRecords(&db, shopID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)

	fmt.Println("SUCCESS")
}

//GetItemByItemID returns item with specific ItemID
func GetItemByItemID(w http.ResponseWriter, r *http.Request) {
	itemID := r.URL.Query().Get("itemid")

	var item []map[string]interface{}
	item = GetSpecificItemRecords(&db, itemID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//Edit User
func EditUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")
	password := r.URL.Query().Get("password")

	var user []map[string]interface{}
	user = EditUserRecords(&db, userID, password)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
}

//Edit Shop
func EditShop(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Query().Get("shopid")
	shopName := r.URL.Query().Get("shopname")
	shopAddress := r.URL.Query().Get("shopaddress")
	shopRating := r.URL.Query().Get("shoprating")
	shopStart := r.URL.Query().Get("shopstart")
	shopEnd := r.URL.Query().Get("shopend")
	shopPostCode := r.URL.Query().Get("shoppostcode")

	var shop []map[string]interface{}
	shop = EditShopRecords(&db, shopID, shopName, shopAddress, shopRating, shopStart, shopEnd, shopPostCode)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)

	fmt.Println("SUCCESS")
}

//Edit Item
func EditItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.URL.Query().Get("itemid")
	itemName := r.URL.Query().Get("itemname")
	itemPrice := r.URL.Query().Get("itemprice")
	itemDesc := r.URL.Query().Get("itemdesc")
	itemImg := r.URL.Query().Get("itemimg")
	shopID := r.URL.Query().Get("shopid")

	var item []map[string]interface{}
	item = EditItemRecords(&db, itemID, itemName, itemPrice, itemDesc, itemImg, shopID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")

	successMessage := DeleteUserRecords(&db, userID)
	returnVal := struct {
		Message string `json:"message"`
		Value   string `json:"value"`
	}{Message: "MessageType", Value: successMessage}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnVal)

	fmt.Println("DELETED")
}

//Delete Shop
func DeleteShop(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Query().Get("shop")

	successMessage := DeleteShopRecords(&db, shopID)
	returnVal := struct {
		Message string `json:"message"`
		Value   string `json:"value"`
	}{Message: "MessageType", Value: successMessage}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnVal)

	fmt.Println("DELETED")
}

//Delete Item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.URL.Query().Get("item")

	successMessage := DeleteItemRecords(&db, itemID)
	returnVal := struct {
		Message string `json:"message"`
		Value   string `json:"value"`
	}{Message: "MessageType", Value: successMessage}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnVal)

	fmt.Println("DELETED")
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

// Middleware function, which will be called for each request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param1 := r.URL.Query().Get("key")

		if param1 != "KEYVALUE"{
			http.Error(w, "Forbidden", http.StatusForbidden)
		}else{
			next.ServeHTTP(w, r)
		}
	})
}