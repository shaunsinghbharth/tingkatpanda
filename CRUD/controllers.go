package CRUD

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func GetUserItems(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")

	var item []map[string]interface{}
	item = GetUserItemsRecords(&db, userID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

// Middleware function, which will be called for each request
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param1 := r.URL.Query().Get("key")

		if param1 != "KEYVALUE"{
			http.Error(w, "Forbidden", http.StatusForbidden)
		}else{
			next.ServeHTTP(w, r)
		}
	})
}
