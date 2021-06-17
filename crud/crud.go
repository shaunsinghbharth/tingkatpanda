package crud

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

var db sql.DB

func Initialise(wg sync.WaitGroup) {
	db = Connect()

	router := mux.NewRouter().StrictSlash(true)

	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":5555", router))

	wg.Done()
}

func initaliseHandlers(router *mux.Router) {
	router.Use(Authenticator)

	router.HandleFunc("/getfullitem", GetFullItem).Methods("GET")
	router.HandleFunc("/getfullitems", GetFullItems).Methods("GET")

	router.HandleFunc("/getuseritems", GetUserItems).Methods("GET")

	router.HandleFunc("/getusers", GetAllUser).Methods("GET")
	router.HandleFunc("/getspecificusers", GetUserByUserName).Methods("GET")
	router.HandleFunc("/editusers", EditUser).Methods("GET")
	router.HandleFunc("/deleteusers", DeleteUser).Methods("GET")

	router.HandleFunc("/getshops", GetAllShop).Methods("GET")
	router.HandleFunc("/getspecificshops", GetShopByShopID).Methods("GET")
	router.HandleFunc("/editshops", EditShop).Methods("GET")
	router.HandleFunc("/deleteshops", DeleteShop).Methods("GET")
	router.HandleFunc("/createshops", CreateShop).Methods("GET")

	router.HandleFunc("/getitems", GetAllItem).Methods("GET")
	router.HandleFunc("/getspecificitems", GetItemByItemID).Methods("GET")
	router.HandleFunc("/edititems", EditItem).Methods("GET")
	router.HandleFunc("/deleteitems", DeleteItem).Methods("GET")
	router.HandleFunc("/createitems", CreateItem).Methods("GET")
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
