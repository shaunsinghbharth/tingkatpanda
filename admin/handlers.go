package admin

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"tingkatpanda/fetcher"
	"tingkatpanda/models"
)

func ServeHTTP(res http.ResponseWriter, req *http.Request){
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		//http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	fmt.Println("Main Handler")
	fmt.Println(req.Cookie("tingkatpanda"))

	title := req.URL.Path[len("/"):]
	fmt.Println(title)
	var p *Page

	//res.Header().Set("Content-Type", "text/css")
	p, err := loadPage(title)

	if err != nil{
		//res.Header().Set("Content-Type", "text/html")
		p, _ = loadPage("admin.html")
	}

	RenderHeader(res,req)

	p.Body.Execute(res, empty{})

	/*
	if manager.IsAdmin(req){
		fmt.Fprint(res, "<a href='/deletesessions/'>Delete All Sessions</a><br>")
		fmt.Fprint(res, "<a href='/deleteusers/'>Delete All Users</a><br>")

	}
	*/

}

func ServeFunctions(res http.ResponseWriter, req *http.Request){

	var p *Page

	p, _ = loadPage("adminfunctions.html")

	RenderHeader(res,req)

	p.Body.Execute(res, empty{})
}

func DeleteItems(res http.ResponseWriter, req *http.Request){
	var itemID string
	vars := mux.Vars(req)
	itemID = vars["id"]
	fmt.Println("ITEMS HANDLER")

	fmt.Println("EDITING ITEM ", itemID)
	//(key string, flag string, itemID string, itemName string, itemPrice string, itemTiming string, itemDesc string, itemImg string, itemCategory string, shopID string)
	fetcher.DeleteItem("KEYVALUE", itemID)

	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		//http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()
	RenderHeader(res,req)

	items := fetcher.GetAllCombinedItem("KEYVALUE")
	var funcMap = template.FuncMap{
		"mod": mod,
		"equal": equal,
		"equalstring": equalstring,
	}

	var tmp = template.Must(template.New("edititem.gohtml").Funcs(funcMap).ParseFiles("htdocs/edititem.gohtml"))
	tmp.Execute(res, &items)
}

func EditItems(res http.ResponseWriter, req *http.Request){
	var itemID string
	var itemName string
	var itemDescription string
	var itemPrice string
	var itemImage string
	var shopID string
	var itemCategory string
	var itemTiming string

	vars := mux.Vars(req)
	itemID = vars["id"]

	fmt.Println("ITEMS HANDLER")

	switch req.Method {
	case "POST":
		req.ParseMultipartForm(10 << 20)

		if req.FormValue("fileUpload") != "" {
			file, handler, err := req.FormFile("fileUpload")
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("File Size: %+v\n", handler.Size)
			fmt.Printf("MIME Header: %+v\n", handler.Header)

			// Create a temporary file within our temp-images directory that follows
			// a particular naming pattern
			tempFile, err := os.Create("htdocs/images/" + handler.Filename)
			if err != nil {
				fmt.Println(err)
			}
			defer tempFile.Close()

			_, err = io.Copy(tempFile, file)
			itemImage = "/images/"+ handler.Filename
		}else{
			itemImage = "/images/temp.png"
		}

		req.ParseForm()
		shopID = req.FormValue("ShopID")
		//itemID = req.FormValue("ItemID")
		itemName = req.FormValue("ItemName")
		itemDescription = req.FormValue("ItemDesc")
		itemPrice = req.FormValue("ItemPrice")
		itemCategory = req.FormValue("ItemCategory")
		itemTiming = req.FormValue("ItemTiming")

		fmt.Println("EDITING ITEM ", itemID)
		//(key string, flag string, itemID string, itemName string, itemPrice string, itemTiming string, itemDesc string, itemImg string, itemCategory string, shopID string)
		fetcher.EditItem("KEYVALUE", "flag", itemID, itemName, itemPrice, itemTiming, itemDescription, itemImage, itemCategory, shopID)
	}

	req.ParseForm()

	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		//http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()
	RenderHeader(res,req)

	items := fetcher.GetAllCombinedItem("KEYVALUE")
	var funcMap = template.FuncMap{
		"mod": mod,
		"equal": equal,
		"equalstring": equalstring,
	}

	var tmp = template.Must(template.New("edititem.gohtml").Funcs(funcMap).ParseFiles("htdocs/edititem.gohtml"))
	tmp.Execute(res, &items)
}

func ServeItems(res http.ResponseWriter, req *http.Request){
	var flag string
	var itemID string
	var itemName string
	var itemDescription string
	var itemPrice string
	var itemImage string
	var shopID string
	var itemCategory string
	var itemTiming string
	var deleteFlag string

	fmt.Println("ITEMS HANDLER")

	switch req.Method {
	case "POST":
		req.ParseMultipartForm(10 << 20)

		file, handler, err := req.FormFile("fileUpload")

			log.Println("no file")
			itemImage = "/images/placeholder.jpg"

			log.Println(err)
			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("File Size: %+v\n", handler.Size)
			fmt.Printf("MIME Header: %+v\n", handler.Header)

			// Create a temporary file within our temp-images directory that follows
			// a particular naming pattern
			defer file.Close()
			tempFile, err := os.Create("htdocs/images/" + handler.Filename)
			if err != nil {
				fmt.Println(err)
			}

			defer tempFile.Close()

			_, err = io.Copy(tempFile, file)
			itemImage = "/images/" + handler.Filename

		defer file.Close()

		req.ParseForm()
		flag = req.FormValue("ITEM")
		shopID = req.FormValue("ShopID")
		itemID = req.FormValue("ItemID")
		itemName = req.FormValue("ItemName")
		itemDescription = req.FormValue("ItemDesc")
		itemPrice = req.FormValue("ItemPrice")
		itemCategory = req.FormValue("ItemCategory")
		itemTiming = req.FormValue("ItemTiming")
		deleteFlag = req.FormValue("deleteFlag")

		fmt.Println("VAL ", flag, itemID, itemName, itemDescription, itemPrice, itemImage, shopID, itemCategory, itemTiming, deleteFlag)
	}

	if deleteFlag == "DELETE"{
		fmt.Println("DELETING ITEM")
		fetcher.DeleteItem("KEYVALUE", itemID)
	}else if deleteFlag == "CREATE"{
		fmt.Println("CREATING SHOP")
		fetcher.CreateItem("KEYVALUE", flag, itemID, itemName, itemPrice, itemTiming, itemDescription, itemImage, itemCategory, shopID)
	} else if flag == "ITEMEDIT"{
		fmt.Println("EDITING ITEM ", itemID)
		//(key string, flag string, itemID string, itemName string, itemPrice string, itemTiming string, itemDesc string, itemImg string, itemCategory string, shopID string)
		fetcher.EditItem("KEYVALUE", flag, itemID, itemName, itemPrice, itemTiming, itemDescription, itemImage, itemCategory, shopID)
	}

	req.ParseForm()

	RenderHeader(res,req)

	items := fetcher.GetAllCombinedItem("KEYVALUE")
	var funcMap = template.FuncMap{
		"mod": mod,
		"equal": equal,
		"equalstring": equalstring,
		"shopIDs" : getShopIDs,
		"shopNames" : getShopNames,
	}

	var tmp = template.Must(template.New("adminitems.gohtml").Funcs(funcMap).ParseFiles("htdocs/adminitems.gohtml"))
	tmp.Execute(res, &items)
}

func getShopIDs() []string{
	shops := fetcher.GetEditShops("KEYVALUE")
	var retVal []string
	for _,v := range shops{
		retVal = append(retVal, v.ShopID)
	}

	return retVal
}

func getShopNames() []models.ShopType{
	shops := fetcher.GetEditShops("KEYVALUE")
	var retVal []models.ShopType
	for _,v := range shops{
		var temp models.ShopType
		temp.ShopName = v.ShopName
		temp.ShopID = v.ShopID
		retVal = append(retVal, temp)
	}

	return retVal
}

func Authenticate(res http.ResponseWriter, req *http.Request){
	fmt.Println("Authentication Redirector")
	fmt.Println(req.Cookie("tingkatpanda"))

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}

	switch req.Method {
	case "POST":
		username := req.FormValue("user")
		fmt.Println("USER ", username)
		password := req.FormValue("password")

		token := manager.CreateSession(username, password)

		fmt.Println(token)
		http.SetCookie(res, token)

		http.Redirect(res,req,"/admin/functions/",http.StatusTemporaryRedirect)
	}
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/admin/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	RenderHeader(res,req)

	p, _ = loadPage("admin.html")

	p.Body.Execute(res, nil)
}

func ServeShops(res http.ResponseWriter, req *http.Request){
	var flag string
	var shopID string
	var shopName string
	var shopAddress string
	var shopRating string
	var shopPostCode string
	var deleteFlag string

	//req.ParseForm()
	switch req.Method {
	case "POST":
		flag = req.FormValue("SHOP")
		shopID = req.FormValue("ShopID")
		shopName = req.FormValue("ShopName")
		shopAddress = req.FormValue("ShopAddress")
		shopRating = req.FormValue("ShopRating")
		shopPostCode = req.FormValue("ShopPostCode")
		deleteFlag = req.FormValue("deleteFlag")

		fmt.Println("deleteFlag", deleteFlag)
	}

	if deleteFlag == "DELETE"{
		fmt.Println("DELETING SHOP")
		fetcher.DeleteShop("KEYVALUE", shopID)
	}else if deleteFlag == "CREATE"{
		fmt.Println("CREATING SHOP")
		fetcher.CreateShop("KEYVALUE", flag, shopID, shopName, shopAddress, shopRating, shopPostCode)
	} else if flag == "SHOPEDIT"{
		fmt.Println("EDITING SHOP ", shopID)
		fetcher.EditShop("KEYVALUE", flag, shopID, shopName, shopAddress, shopRating, shopPostCode)
	}

	var p *Page

	p, _ = loadPage("adminshops.gohtml")

	RenderHeader(res,req)

	var shops []models.ShopsEdit
	shops = fetcher.GetEditShops("KEYVALUE")
	p.Body.Execute(res, shops)
}

func ServeUsers(res http.ResponseWriter, req *http.Request){
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/admin/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	var p *Page

	p, _ = loadPage("adminusers.gohtml")

	RenderHeader(res,req)

	p.Body.Execute(res, empty{})
}

func Login(res http.ResponseWriter, req *http.Request){
	fmt.Println("Login Page")
	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	http.Redirect(res, req,"/admin/functions/", http.StatusTemporaryRedirect)

	p, _ = loadPage("login.html")

	//RenderHeader(res,req)
	p.Body.Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
	p.Body.Execute(res, nil)
}

func RenderHeader(res http.ResponseWriter, req *http.Request){

	var funcMap = template.FuncMap{
		"mod": mod,
		"equal": equal,
	}

	tmp := template.Must(template.New("header.gohtml").Funcs(funcMap).ParseFiles("htdocs/header.gohtml"))
	fmt.Println("SESSIONUSER ", manager.GetCurrentUserObject(req))
	tmp.Execute(res, manager.GetCurrentUserObject(req))
}

func mod(i, j int) bool {
	return i%j == 0
}

func equal(i, j int) bool {
	return i == j
}

func equalstring(i int, j string) bool {
	value, _ := strconv.Atoi(j)
	return i == value
}

func loginHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("Login Handler")
	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	p = p

	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}

	switch req.Method{
	case "POST":
		username := req.FormValue("username")
		password := req.FormValue("password")

		token := manager.CreateSession(username,password)

		fmt.Println(token)
		http.SetCookie(res, token)

		if token == nil{
			p, err := loadPage("login.gohtml")
			if err != nil{
				fmt.Println("Error")
			}
			p.Body.Execute(res,empty{})

			fmt.Fprintf(res, "Login Credentials are incorrect. Did not log in.")
		}else{
			http.Redirect(res,req,"/",http.StatusTemporaryRedirect)
		}
		return

	default:
	}
}

func Destroy(res http.ResponseWriter, req *http.Request){
	token := manager.DestroySession(req)
	http.SetCookie(res, token)

	http.Redirect(res,req,"/",http.StatusTemporaryRedirect)
}