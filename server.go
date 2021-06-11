package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct { // map this type to the record in the Users table
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Shops struct { // map this type to the record in the Shops table
	ShopId      int    `json:"shopid"`
	ShopName    string `json:"shopname"`
	ShopAddress string `json:"shopaddress"`
	ShopRating  string `json:"shoprating"`
	ShopPeriod  string `json:"shopperiod"`
}

type Items struct { // map this type to the record in the Items table
	ItemId    int     `json:"itemid"`
	ItemName  string  `json:"itemname"`
	ItemPrice float64 `json:"itemprice"`
	ItemDesc  string  `json:"itemdesc"`
	ItemImg   string  `json:"itemimg"`
	ShopId    int     `json:"shopid"`
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
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	var user []Users
	user = GetUserRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
}

//GetAllShop get all shop data
func GetAllShop(w http.ResponseWriter, r *http.Request) {

	var shop []Shops
	shop = GetShopRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)

	fmt.Println("SUCCESS")
}

//GetAllItem get all item data
func GetAllItem(w http.ResponseWriter, r *http.Request) {

	var item []Items
	item = GetItemRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//GetUserByUserName returns user with specific UserName
func GetUserByUserName(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("user")

	var user []Users
	user = GetSpecificUserRecords(&db, param1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
}

//GetShopByShopID returns shop with specific ShopID
func GetShopByShopID(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("shopid")

	var shop []Shops
	shop = GetSpecificShopRecords(&db, param1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)

	fmt.Println("SUCCESS")
}

//GetItemByItemID returns item with specific ItemID
func GetItemByItemID(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("itemid")

	var item []Items
	item = GetSpecificItemRecords(&db, param1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//Edit User
func EditUser(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("user")
	param2 := r.URL.Query().Get("password")

	var user []Users
	user = EditUserRecords(&db, param1, param2)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
}

//Edit Shop
func EditShop(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("shopid")
	param2 := r.URL.Query().Get("shopname")
	param3 := r.URL.Query().Get("shopaddress")
	param4 := r.URL.Query().Get("shoprating")
	param5 := r.URL.Query().Get("shopperiod")

	var shop []Shops
	shop = EditShopRecords(&db, param1, param2, param3, param4, param5)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)

	fmt.Println("SUCCESS")
}

//Edit Item
func EditItem(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("itemid")
	param2 := r.URL.Query().Get("itemname")
	param3 := r.URL.Query().Get("itemprice")
	param4 := r.URL.Query().Get("itemdesc")
	param5 := r.URL.Query().Get("itemimg")
	param6 := r.URL.Query().Get("shopid")

	var item []Items
	item = EditItemRecords(&db, param1, param2, param3, param4, param5, param6)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("user")

	successMessage := DeleteUserRecords(&db, param1)
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
	param1 := r.URL.Query().Get("shop")

	successMessage := DeleteShopRecords(&db, param1)
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
	param1 := r.URL.Query().Get("item")

	successMessage := DeleteItemRecords(&db, param1)
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

func GetUserRecords(db *sql.DB) []Users {
	results, err := db.Query("Select * FROM GOLIVEDB.Users")

	if err != nil {
		panic(err.Error())
	}

	var returnVal []Users

	for results.Next() {
		var user Users
		results.Scan(&user.UserName, &user.Password)
		returnVal = append(returnVal, user)
	}

	return returnVal
}

func GetShopRecords(db *sql.DB) []Shops {
	results, err := db.Query("Select * FROM GOLIVEDB.Shops")

	if err != nil {
		panic(err.Error())
	}

	var returnVal []Shops

	for results.Next() {
		var shop Shops
		results.Scan(&shop.ShopId, &shop.ShopName, &shop.ShopAddress, &shop.ShopRating, &shop.ShopPeriod)
		returnVal = append(returnVal, shop)
	}

	return returnVal
}

func GetItemRecords(db *sql.DB) []Items {
	results, err := db.Query("Select * FROM GOLIVEDB.Items")

	if err != nil {
		panic(err.Error())
	}

	var returnVal []Items

	for results.Next() {
		var item Items
		results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)
		returnVal = append(returnVal, item)
	}

	return returnVal
}

func GetSpecificUserRecords(db *sql.DB, AN string) []Users {
	results, err := db.Query("Select * FROM GOLIVEDB.Users WHERE UserName=?", AN)

	if err != nil {
		panic(err.Error())
	}

	var returnVal []Users

	for results.Next() {
		var user Users
		results.Scan(&user.UserName, &user.Password)
		returnVal = append(returnVal, user)
	}

	return returnVal
}

func GetSpecificShopRecords(db *sql.DB, SI string) []Shops {
	results, err := db.Query("Select * FROM GOLIVEDB.Shops WHERE ShopID=?", SI)

	if err != nil {
		panic(err.Error())
	}

	var returnVal []Shops

	for results.Next() {
		var shop Shops
		results.Scan(&shop.ShopId, &shop.ShopName, &shop.ShopAddress, &shop.ShopRating, &shop.ShopPeriod)
		returnVal = append(returnVal, shop)
	}

	return returnVal
}

func GetSpecificItemRecords(db *sql.DB, II string) []Items {
	results, err := db.Query("Select * FROM GOLIVEDB.Items WHERE ItemID=?", II)

	if err != nil {
		panic(err.Error())
	}

	var returnVal []Items

	for results.Next() {
		var item Items
		results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)
		returnVal = append(returnVal, item)
	}

	return returnVal
}

func EditUserRecords(db *sql.DB, AN string, AK string) []Users {
	results, err := db.Query("UPDATE GOLIVEDB.Users SET Password=? WHERE UserName=?", AK, AN)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Users WHERE UserName=?", AN)
	}

	var returnVal []Users

	for results.Next() {
		var user Users
		results.Scan(&user.UserName, &user.Password)
		returnVal = append(returnVal, user)
	}

	return returnVal
}

func EditShopRecords(db *sql.DB, ID string, SN string, SA string, SR string, SP string) []Shops {
	//func EditShopRecords(db *sql.DB, ID int, SN string, SA string, SR string, SP string) []Shops {
	results, err := db.Query("UPDATE Shops SET ShopName=?, ShopAddress=?, ShopRating=?, ShopPeriod=? WHERE ShopId=?", SN, SA, SR, SP, ID)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Shops WHERE ShopID=?", ID)
	}

	var returnVal []Shops

	for results.Next() {
		var shop Shops
		results.Scan(&shop.ShopId, &shop.ShopName, &shop.ShopAddress, &shop.ShopRating, &shop.ShopPeriod)
		returnVal = append(returnVal, shop)
	}

	return returnVal
}

func EditItemRecords(db *sql.DB, ID string, IN string, IP string, DE string, IG string, SI string) []Items {
	//func EditItemRecords(db *sql.DB, ID int, IN string, IP float64, DE string, IG string, SI int) []Items {
	results, err := db.Query("UPDATE Items SET ItemName=?, ItemPrice=?, ItemDesc=?, ItemImg=?, ShopID=? WHERE ItemId=?", IN, IP, DE, IG, SI, ID)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Items WHERE ItemID=?", ID)
	}

	var returnVal []Items

	for results.Next() {
		var item Items
		results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)
		returnVal = append(returnVal, item)
	}

	return returnVal
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
