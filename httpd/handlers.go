package httpd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"sync"
	"tingkatpanda/enginator"
	"tingkatpanda/fetcher"
	"tingkatpanda/models"
)

func ServeHTTP(res http.ResponseWriter, req *http.Request){
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
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
		p, _ = loadPage("index.html")
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

func ShowSelect(res http.ResponseWriter, req *http.Request){
	postcode := req.URL.Query().Get("postcode")

	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	RenderHeader(res,req)

	p, _ = loadPage("select.html")

	p.Body.Execute(res, postcode)
}

func Populate(key string) *enginator.EnginatorTable{
	table := enginator.Table("recommendations")

	users := fetcher.GetUsers(key)

	for _,v := range users{
		temp := make(map[interface{}]float64)
		items := fetcher.GetUserItems(key,v.UserName)

		for _,val := range items{
			fmt.Println(v.UserName)
			rating := float64(val.Rating)
			temp[val.ItemID] = rating

			fmt.Println("POPULATING ", val.ItemID, val.Rating)
		}
		table.Add(v.UserName, temp)
	}
	return table
}

func GetCoordinates(postcode string) (float64, float64)  {
	///commonapi/search?searchVal={SearchText}&returnGeom={Y/N}&getAddrDetails={Y/N}&pageNum={PageNumber}
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://developers.onemap.sg/commonapi/search", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("searchVal", postcode)
	q.Add("returnGeom", "Y")
	q.Add("getAddrDetails", "Y")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(string(bodyBytes))

	var data map[string]interface{}
	//err = json.NewDecoder(resp.Body).Decode(&data)
	json.Unmarshal(bodyBytes, &data)

	if err != nil {
		//Error
	}

	var results []interface{}
	var singleResult map[string]interface{}
	for k,v := range data{
		fmt.Println("COORDSMAP", k,v)
		if k == "results"{
			results = v.([]interface{})
			singleResult = results[0].(map[string]interface{})
		}
	}

	fmt.Println("RESULT" , singleResult)
	//return 0.0, 0.0
	fmt.Println("SINGLERESULT ", singleResult)
	LAT, _ := strconv.ParseFloat(singleResult["LATITUDE"].(string), 64)
	LONG, _ := strconv.ParseFloat(singleResult["LONGITUDE"].(string), 64)
	return LAT, LONG
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1) * math.Sin(radlat2) + math.Cos(radlat1) * math.Cos(radlat2) * math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
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

		http.Redirect(res,req,"/",http.StatusTemporaryRedirect)
	}
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	RenderHeader(res,req)

	p, _ = loadPage("login.html")

	p.Body.Execute(res, nil)
}

func Login(res http.ResponseWriter, req *http.Request){
	fmt.Println("Login Page")
	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	p, _ = loadPage("login.html")

	RenderHeader(res,req)
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

func ShowRecommendation(res http.ResponseWriter, req *http.Request){
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()


	var user string
	var postcode string
	var timings []string
	var category string
	var price float64

	switch req.Method {
	case "POST":
		user = req.FormValue("user")
		postcode = req.FormValue("postcode")
		timings = req.Form["timing"]
		category = req.FormValue("category")
		price, _ = strconv.ParseFloat(req.FormValue("price"),64)

		fmt.Println("POSTTIMINGS ", timings)
	}
	fmt.Println("POSTPOSTCODE: ", postcode)
	lat, long := GetCoordinates(postcode)

	table := Populate("KEYVALUE")
	recs, _ := table.Recommend(user)

	var output map[int][]models.CombinedItem
	output = make(map[int][]models.CombinedItem)
	for i, rec := range recs {
		fmt.Println("Recommending", rec.Key, "with score:", rec.Distance, "at index:", i)

		var tempItem []models.CombinedItem
		//cacheOutput, found := c.Get(strconv.Itoa(rec.Key.(int)))
		//if found {
		//	tempItem = cacheOutput.([]models.CombinedItem)
		//} else {
			tempItem = fetcher.GetCombinedItem("KEYVALUE", strconv.Itoa(rec.Key.(int)))
			fmt.Println("COMBINEDITEMS ", tempItem)
			mutex.Lock()
			//c.Set(strconv.Itoa(rec.Key.(int)), output, cachinator.DefaultExpiration)
			mutex.Unlock()
		//}

		fmt.Println("tempitempost ", tempItem[0])
		shoplat, shoplong := GetCoordinates(tempItem[0].ShopPostCode)
		itemPrice, _ := strconv.ParseFloat(tempItem[0].ItemPrice, 64)
		itemCategory := tempItem[0].ItemCategory
		dist := distance(lat,long,shoplat,shoplong, "K")
		if dist < 5 && price == itemPrice && category == itemCategory{
			for _,v := range timings {
				itemTiming, _ := strconv.Atoi(tempItem[0].ItemTiming)
				fmt.Println("TIMING ", itemTiming)
				idx, _ := strconv.Atoi(v)
				if v == "0" {
					output[idx] = append(output[idx], tempItem[0])
				}
				if v == "1" {
					output[idx] = append(output[idx], tempItem[0])
				}
				if v == "2" {
					output[idx] = append(output[idx], tempItem[0])
				}
			}
		}
	}

	for i,_ := range output{
		if len(output[i]) > 7 {
			output[i] = output[i][:7]
		}
	}

	RenderHeader(res,req)

	var funcMap = template.FuncMap{
		"mod": mod,
		"equal": equal,
	}

	var tmp = template.Must(template.New("results.gohtml").Funcs(funcMap).ParseFiles("htdocs/results.gohtml"))
	err := tmp.Execute(res, &output)

	fmt.Println("ERROR ", err)
}

func mod(i, j int) bool {
	return i%j == 0
}

func equal(i, j int) bool {
	return i == j
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

func signupHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("Signup Handler")

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
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")

		token := manager.RegisterUser(username,password,firstname,lastname)

		fmt.Println(token)
		http.SetCookie(res, token)

		if token == nil{
			fmt.Fprintf(res, "Login Credentials are incorrect. Did not log in.")
		}else{
			fmt.Fprintf(res, "Name = %s\n", username)
			fmt.Fprintf(res, "Address = %s\n", password)
			fmt.Fprintf(res, "Token = %s\n", token)
		}
	default:
	}
}


func adminHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("Main Handler")

	title := req.URL.Path[len("/"):]
	p, _ := loadPage(title)
	p.Body.Execute(res, empty{})
}

func profileEditor(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case "POST":
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")

		uuid := manager.GetCurrentUserObject(req).UUID
		manager.EditUser(uuid,password, username, firstname,lastname )
		manager.Dump()
		break
	default:

		}

	p, _ := loadPage("profile.gohtml")
	data := manager.GetCurrentUserObject(req)

	p.Body.Execute(res, data)
}

func Destroy(res http.ResponseWriter, req *http.Request){
	token := manager.DestroySession(req)
	http.SetCookie(res, token)

	http.Redirect(res,req,"/",http.StatusTemporaryRedirect)
}

func DELETEALLSESSIONS(res http.ResponseWriter, req *http.Request) {
	fmt.Println("DESTROYING ALL SESSIONS")
	io.WriteString(res, "<html><script type=\"text/javascript\">\nalert(\"DESTROYING ALL SESSIONS\");</script>\n</html>")

	for _,v := range manager.Users{
		fmt.Println("USER DELETING ", v)
		manager.DeleteSession(v.Username)
	}
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}

func DELETEALLUSERS(res http.ResponseWriter, req *http.Request) {
	fmt.Println("DESTORYING ALL USERS AND FREEING UP BOOKINGS")
	io.WriteString(res, "<html><script type=\"text/javascript\">alert(\"DESTROYING ALL USERS\");</script>\n</html>")

	for _,v := range manager.Users{
		manager.DeleteUser(v.Username)
	}
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}