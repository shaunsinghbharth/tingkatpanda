package crud

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
}

func GetFullItems(w http.ResponseWriter, r *http.Request) {

	var item []map[string]interface{}
	item = GetFullItemRecordsAll(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

//GetAllUser get all user data
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	var user []map[string]interface{}
	user = GetUserRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//GetAllShop get all shop data
func GetAllShop(w http.ResponseWriter, r *http.Request) {

	var shop []map[string]interface{}
	shop = GetShopRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

//GetAllItem get all item data
func GetAllItem(w http.ResponseWriter, r *http.Request) {

	var item []map[string]interface{}
	item = GetItemRecords(&db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

//GetUserByUserName returns user with specific UserName
func GetUserByUserName(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")

	var user []map[string]interface{}
	user = GetSpecificUserRecords(&db, userID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//GetShopByShopID returns shop with specific ShopID
func GetShopByShopID(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Query().Get("shopid")

	var shop []map[string]interface{}
	shop = GetSpecificShopRecords(&db, shopID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

//GetItemByItemID returns item with specific ItemID
func GetItemByItemID(w http.ResponseWriter, r *http.Request) {
	itemID := r.URL.Query().Get("itemid")

	var item []map[string]interface{}
	item = GetSpecificItemRecords(&db, itemID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
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
}

//Edit Shop
func EditShop(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Query().Get("shopID")
	shopName := r.URL.Query().Get("shopName")
	shopAddress := r.URL.Query().Get("shopAddress")
	shopRating := r.URL.Query().Get("shopRating")
	shopStart := r.URL.Query().Get("shopstart")
	shopEnd := r.URL.Query().Get("shopend")
	shopPostCode := r.URL.Query().Get("shopPostCode")

	var shop []map[string]interface{}
	shop = EditShopRecords(&db, shopID, shopName, shopAddress, shopRating, shopStart, shopEnd, shopPostCode)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

//Create Shop
func CreateShop(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Query().Get("shopID")
	shopName := r.URL.Query().Get("shopName")
	shopAddress := r.URL.Query().Get("shopAddress")
	shopRating := r.URL.Query().Get("shopRating")
	shopStart := r.URL.Query().Get("shopstart")
	shopEnd := r.URL.Query().Get("shopend")
	shopPostCode := r.URL.Query().Get("shopPostCode")

	var shop []map[string]interface{}
	shop = CreateShopRecords(&db, shopID, shopName, shopAddress, shopRating, shopStart, shopEnd, shopPostCode)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

//Edit Item
func EditItem(w http.ResponseWriter, r *http.Request) {

	itemID := r.URL.Query().Get("itemID")
	itemName := r.URL.Query().Get("itemName")
	itemPrice := r.URL.Query().Get("itemPrice")
	itemDesc := r.URL.Query().Get("itemDesc")
	itemImg := r.URL.Query().Get("itemImg")
	itemCategory := r.URL.Query().Get("itemCategory")
	shopID := r.URL.Query().Get("shopID")
	itemTiming := r.URL.Query().Get("itemTiming")

	fmt.Println("ITEMID Q ", itemID)

	var item []map[string]interface{}
	item = EditItemRecords(&db, itemID, itemName, itemCategory, itemPrice, itemDesc, itemImg,itemTiming, shopID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}


//Edit Item
func CreateItem(w http.ResponseWriter, r *http.Request) {

	itemID := r.URL.Query().Get("itemID")
	itemName := r.URL.Query().Get("itemName")
	itemPrice := r.URL.Query().Get("itemPrice")
	itemDesc := r.URL.Query().Get("itemDesc")
	itemImg := r.URL.Query().Get("itemImg")
	itemCategory := r.URL.Query().Get("itemCategory")
	shopID := r.URL.Query().Get("shopID")
	itemTiming := r.URL.Query().Get("itemTiming")

	fmt.Println("ITEMID Q ", itemID)

	CreateItemRecords(&db, itemID, itemName, itemCategory, itemPrice, itemDesc, itemImg,itemTiming, shopID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
}

//Delete Shop
func DeleteShop(w http.ResponseWriter, r *http.Request) {
	shopID := r.URL.Query().Get("shopID")

	fmt.Println("DELETING SHOP")
	successMessage := DeleteShops(&db, shopID)
	returnVal := struct {
		Message string `json:"message"`
		Value   string `json:"value"`
	}{Message: "MessageType", Value: successMessage}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnVal)
}

//Delete Item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.URL.Query().Get("itemID")

	successMessage := DeleteItemRecords(&db, itemID)
	returnVal := struct {
		Message string `json:"message"`
		Value   string `json:"value"`
	}{Message: "MessageType", Value: successMessage}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnVal)
}

func GetUserItems(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")

	var item []map[string]interface{}
	item = GetUserItemsRecords(&db, userID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
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
