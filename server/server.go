// Name : Siew Tuck Meng
// Email :tuckmengsiew@gmail.com
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"database/sql"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-sanitize/sanitize"
)

type Users struct { // map this type to the record in the Users table
	UserName string `json:"username" , san:"max=55,trim,lower"`
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

	// use a randomly generated very large integer
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	subject := pkix.Name{
		Organization:       []string{"green Co."},
		OrganizationalUnit: []string{"green"},
		CommonName:         "green",
	}
	// crypto/x509 is used to create the certificate
	// X.509 is a standard defining the format of public key certificates
	// X.509 certificates are used in many Internet protocols, including TLS/SSL, which is the basis for HTTPS,[2] the secure protocol for browsing the web.
	template := x509.Certificate{
		SerialNumber: serialNumber, // Certificate serial number is unique number issued by CA.
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),                         // Validity period set as one year from day the certificate is created.
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // certificate is used for server authentication
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // certificate is used for server authentication
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}
	// create a RSA private key using crypto/rsa by calling GenerateKey()
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// create the cert
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// create the private key
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

	log.Fatal(http.ListenAndServe(":5555", router))
}

func PandaMenu(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Tingkat Panda Backend Server")
}

func initaliseHandlers(router *mux.Router) {

	router.HandleFunc("/", PandaMenu).Methods("GET")

	router.HandleFunc("/registernewuser", InsertUser).Methods("GET")
	router.HandleFunc("/userlogin", UserLogin).Methods("GET")
	router.HandleFunc("/getusers", GetAllUser).Methods("GET")
	router.HandleFunc("/getspecificusers", GetUserByUserName).Methods("GET")
	router.HandleFunc("/editusers", EditUser).Methods("GET")
	router.HandleFunc("/deleteusers", DeleteUser).Methods("GET")

	router.HandleFunc("/insertshops", InsertShop).Methods("GET")
	router.HandleFunc("/getshops", GetAllShop).Methods("GET")
	router.HandleFunc("/getspecificshops", GetShopByShopID).Methods("GET")
	router.HandleFunc("/editshops", EditShop).Methods("GET")
	router.HandleFunc("/deleteshops", DeleteShop).Methods("GET")

	router.HandleFunc("/insertitems", InsertItem).Methods("GET")
	router.HandleFunc("/getitems", GetAllItem).Methods("GET")
	router.HandleFunc("/getspecificitems", GetItemByItemID).Methods("GET")
	router.HandleFunc("/edititems", EditItem).Methods("GET")
	router.HandleFunc("/deleteitems", DeleteItem).Methods("GET")

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
	if err != nil {
		return false
	}
	return true
}

//Insert User
func InsertUser(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("user")
	param2 := r.URL.Query().Get("password")

	var user []Users

	//insert the (trim) username and (hash) password into the DB
	user = InsertUserRecords(&db, param1, string(HashPassword(param2)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
}

//Insert Shop
func InsertShop(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("shopid")
	param2 := r.URL.Query().Get("shopname")
	param3 := r.URL.Query().Get("shopaddress")
	param4 := r.URL.Query().Get("shoprating")
	param5 := r.URL.Query().Get("shopperiod")

	var shop []Shops
	shop = InsertShopRecords(&db, param1, param2, param3, param4, param5)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)

	fmt.Println("SUCCESS")
}

//Insert Item
func InsertItem(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("itemid")
	param2 := r.URL.Query().Get("itemname")
	param3 := r.URL.Query().Get("itemprice")
	param4 := r.URL.Query().Get("itemdesc")
	param5 := r.URL.Query().Get("itemimg")
	param6 := r.URL.Query().Get("shopid")

	var item []Items
	item = InsertItemRecords(&db, param1, param2, param3, param4, param5, param6)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("SUCCESS")
}

//User Login
func UserLogin(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("user")
	param2 := r.URL.Query().Get("password")

	var user []Users
	user = GetHashedPasswordRecords(&db, param1, param2)

	// the password saved in the db the user's supplied password
	if VerifyPassword([]byte(param1), param2) {
		fmt.Println("User authenticated!")
	} else {
		fmt.Println("Invalid username and/or password")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	fmt.Println("SUCCESS")
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

	//edit the (trim) username and (hash) password into the DB
	user = EditUserRecords(&db, param1, string(HashPassword(param2)))

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
	param1 := r.URL.Query().Get("shopid")

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
	param1 := r.URL.Query().Get("itemid")

	successMessage, err := DeleteItemRecords(&db, param1)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		returnVal := struct {
			Message string `json:"message"`
			Value   string `json:"value"`
		}{Message: "MessageType", Value: successMessage}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		s, _ := sanitize.New()
		s.Sanitize(&returnVal)

		json.NewEncoder(w).Encode(returnVal)

		fmt.Println("DELETED")
	}
}

func Connect() sql.DB {
	db, err := sql.Open("mysql", "tuckmeng:G0L1V3@tcp(128.199.125.231:3306)/GOLIVEDB")
	//db, err := sql.Open("mysql", "user1:password@tcp(127.0.0.1:3306)/GOLIVEDB")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened.")
	}
	defer db.Close()

	return *db

}

func InsertUserRecords(db *sql.DB, username, password string) []Users {
	results, err := db.Query("INSERT INTO GOLIVEDB.Users VALUES (?,?)", username, password)

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

func InsertShopRecords(db *sql.DB, shopid, shopname, shopaddress, shoprating, shopperiod string) []Shops {
	results, err := db.Query("INSERT INTO GOLIVEDB.Shop VALUES (?,?,?,?,?)", shopid, shopname, shopaddress, shoprating, shopperiod)

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

func InsertItemRecords(db *sql.DB, itemid, itemname, itemprice, itemdesc, itemimg, shopid string) []Items {
	results, err := db.Query("INSERT INTO GOLIVEDB.Item VALUES (?,?,?,?,?,?)", itemid, itemname, itemprice, itemdesc, itemimg, shopid)

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
	results, err := db.Query("Select * FROM GOLIVEDB.Shop")

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
	results, err := db.Query("Select * FROM GOLIVEDB.Item")

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

// get the hashed password of the user in string type
func GetHashedPasswordRecords(db *sql.DB, UN, PW string) []Users {
	results, err := db.Query("SELECT * FROM GOLIVEDB.Users where Username=?", UN)

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

func GetSpecificUserRecords(db *sql.DB, UN string) []Users {
	results, err := db.Query("Select * FROM GOLIVEDB.Users WHERE UserName=?", UN)

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
	results, err := db.Query("Select * FROM GOLIVEDB.Shop WHERE ShopID=?", SI)

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
	results, err := db.Query("Select * FROM GOLIVEDB.Item WHERE ItemID=?", II)

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

func EditUserRecords(db *sql.DB, UN string, PW string) []Users {
	results, err := db.Query("UPDATE GOLIVEDB.Users SET Password=? WHERE UserName=?", PW, UN)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Users WHERE UserName=?", UN)
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
	results, err := db.Query("UPDATE Shop SET ShopName=?, ShopAddress=?, ShopRating=?, ShopPeriod=? WHERE ShopId=?", SN, SA, SR, SP, ID)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Shop WHERE ShopID=?", ID)
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
	results, err := db.Query("UPDATE Item SET ItemName=?, ItemPrice=?, ItemDesc=?, ItemImg=?, ShopID=? WHERE ItemId=?", IN, IP, DE, IG, SI, ID)

	if err != nil {
		panic(err.Error())
	} else {
		results, err = db.Query("Select * FROM GOLIVEDB.Item WHERE ItemID=?", ID)
	}

	var returnVal []Items

	for results.Next() {
		var item Items
		results.Scan(&item.ItemId, &item.ItemName, &item.ItemPrice, &item.ItemDesc, &item.ItemImg, &item.ShopId)
		returnVal = append(returnVal, item)
	}

	return returnVal
}

func DeleteUserRecords(db *sql.DB, UN string) string {
	results, err := db.Query("DELETE FROM Users WHERE UserName=?", UN)

	if err != nil {
		return "404 - User Does Not Exist"
		//panic(err.Error())
	}

	if results != nil {
		return "Error Deleting"
	}

	return "201 - Delete Success"
}

func DeleteShopRecords(db *sql.DB, ID string) string {
	results, err := db.Query("DELETE FROM Shop WHERE ShopID=?", ID)

	if err != nil {
		return "404 - Shop Does Not Exist"
		//panic(err.Error())
	}

	if results != nil {
		return "Error Deleting"
	}

	return "201 - Delete Success"
}

func DeleteItemRecords(db *sql.DB, ID string) (string, error) {
	_, err := db.Query("DELETE FROM Item WHERE ItemID=?", ID)

	if err != nil {
		return "", errors.New("does not exist")
		//panic(err.Error())
	}

	return "201 - Delete Success", nil
}
