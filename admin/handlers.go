package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"tingkatpanda/fetcher"
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
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		//http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	var p *Page

	p, _ = loadPage("adminfunctions.html")

	RenderHeader(res,req)

	p.Body.Execute(res, empty{})
}

func ServeItems(res http.ResponseWriter, req *http.Request){
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		//http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	var p *Page

	p, _ = loadPage("adminitems.gohtml")

	RenderHeader(res,req)

	items := fetcher.GetAllCombinedItem("KEYVALUE")
	p.Body.Execute(res, items)
}

func ServeShops(res http.ResponseWriter, req *http.Request){
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		//http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
	}
	mutex.Unlock()

	var p *Page

	p, _ = loadPage("adminshops.gohtml")

	RenderHeader(res,req)

	shops := fetcher.GetShops("KEYVALUE")
	p.Body.Execute(res, shops)
}

func ServeUsers(res http.ResponseWriter, req *http.Request){
	var mutex = &sync.Mutex{}

	mutex.Lock()
	if manager.ValidSession(req) == false{
		//http.Redirect(res,req,"/login/",http.StatusTemporaryRedirect)
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

func Destroy(res http.ResponseWriter, req *http.Request){
	token := manager.DestroySession(req)
	http.SetCookie(res, token)

	http.Redirect(res,req,"/",http.StatusTemporaryRedirect)
}