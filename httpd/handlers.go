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
	"tingkatpanda/enginator"
	"tingkatpanda/models"
)

func ServeHTTP(res http.ResponseWriter, req *http.Request){
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}

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
	p.Body.Execute(res, empty{})

	/*
	if manager.IsAdmin(req){
		fmt.Fprint(res, "<a href='/deletesessions/'>Delete All Sessions</a><br>")
		fmt.Fprint(res, "<a href='/deleteusers/'>Delete All Users</a><br>")

	}
	*/

}

func ShowSelect(res http.ResponseWriter, req *http.Request){
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}

	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	p, _ = loadPage("select.html")

	p.Body.Execute(res, nil)
}

func Populate(key string) *enginator.EnginatorTable{
	table := enginator.Table("recommendations")

	users := enginator.GetUsers(key)

	for _,v := range users{
		temp := make(map[interface{}]float64)
		items := enginator.GetUserItems(key,v.UserName)

		for _,val := range items{
			fmt.Println(v.UserName)
			rating := float64(val.Rating)
			temp[val.ItemID] = rating
		}
		table.Add(v.UserName, temp)
	}

	return table
}

func GetCoordinates(postcode string) (float64, float64)  {
	fmt.Println("COORDSPOSTCODE", postcode)
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

	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}

	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	p, _ = loadPage("login.html")

	p.Body.Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
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

	p.Body.Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
	p.Body.Execute(res, nil)
}

func ShowRecommendation(res http.ResponseWriter, req *http.Request){
	if manager.ValidSession(req) == false{
		http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}

	user := req.URL.Query().Get("user")
	postcode := req.URL.Query().Get("postcode")
	fmt.Println("POSTCODE ", postcode)
	price, _ := strconv.ParseFloat(req.URL.Query().Get("price"), 64)
	category := req.URL.Query().Get("category")
	lat, long := GetCoordinates(postcode)

	table := Populate("KEYVALUE")
	recs, _ := table.Recommend(user)

	var output []models.CombinedItem
		for i, rec := range recs {
			fmt.Println("Recommending", rec.Key, "with score:", rec.Distance, "at index:", i)

				var tempItem []models.CombinedItem
				//cacheOutput, found := c.Get(strconv.Itoa(rec.Key.(int)))
				//if found {
					//temp = cacheOutput.([]models.CombinedItem)
				//} else {
					tempItem = enginator.GetCombinedItem("KEYVALUE", strconv.Itoa(rec.Key.(int)))
					//c.Set(strconv.Itoa(rec.Key.(int)), output, cache.DefaultExpiration)
				//}
				shoplat, shoplong := GetCoordinates(tempItem[0].ShopPostCode)
				itemPrice := tempItem[0].ItemPrice
				itemCategory := tempItem[0].ItemCategory
				dist := distance(lat,long,shoplat,shoplong, "K")
				if dist < 5 && price == itemPrice && category == itemCategory{
					output = append(output, tempItem[0])
				}
		}


	fmt.Println("Recommendation Handler")
	var p *Page

	p = &Page{
		Title: "",
		Body:  nil,
	}

	p, _ = loadPage("results.gohtml")

	fmt.Println("OUTPUT ", output)
	p.Body.Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
	p.Body.Execute(res, &output)
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