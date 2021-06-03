package http

import (
	"assignment3/bookings"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func serveHTTP(res http.ResponseWriter, req *http.Request){
	fmt.Println("Main Handler")
	fmt.Println(req.Cookie("BharthPHD"))

	title := req.URL.Path[len("/"):]
	//fmt.Println(title)
	var p *Page

	//res.Header().Set("Content-Type", "text/css")
	p, err := loadPage(title)

	if err != nil{
		//res.Header().Set("Content-Type", "text/html")
		p, _ = loadPage("index.gohtml")
	}
	p.Body.Execute(res, empty{})

	if manager.IsAdmin(req){
		fmt.Fprint(res, "<a href='/deletesessions/'>Delete All Sessions</a><br>")
		fmt.Fprint(res, "<a href='/deleteusers/'>Delete All Users</a><br>")

	}
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
		p, err := loadPage("login.gohtml")
		if err != nil{
			return
		}
		p.Body.Execute(res,empty{})
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

		p, err := loadPage("signup.gohtml")
		if err != nil{
			return
		}

		token := manager.RegisterUser(username,password,firstname,lastname)

		fmt.Println(token)
		http.SetCookie(res, token)


		p.Body.Execute(res,empty{})

		if token == nil{
			fmt.Fprintf(res, "Login Credentials are incorrect. Did not log in.")
		}else{
			fmt.Fprintf(res, "Name = %s\n", username)
			fmt.Fprintf(res, "Address = %s\n", password)
			fmt.Fprintf(res, "Token = %s\n", token)
		}

	default:
		p, err := loadPage("signup.gohtml")
		if err != nil{
			return
		}
		p.Body.Execute(res,empty{})
	}
}


func adminHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("Main Handler")

	title := req.URL.Path[len("/"):]
	p, _ := loadPage(title)
	p.Body.Execute(res, empty{})
}

func viewHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("View Handler")
	switch req.Method {
	case "GET":
		id := req.FormValue("id")
		if id == "" {
			break
		}
		p, err := loadPage("edit.gohtml")
		if err != nil {
			return
		}

		slot, err := strconv.Atoi(id)
		data := bookings.GenerateCurrentBooking(slot, manager.GetCurrentUser(req))

		p.Body.Execute(res, data)
		return
	case "POST":
		time := req.FormValue("time")
		doctor, _ := strconv.Atoi(req.FormValue("doctor"))
		slot, _ := strconv.Atoi(req.FormValue("slot"))

		bookings.EditBooking(slot,time,doctor)
	}

	p, _ := loadPage("view.gohtml")
	if bookings.ReturnCurrentUserBookings(manager.GetCurrentUser(req)) == nil{
		p.Body.Execute(res, empty{})

	}else{
		p.Body.Execute(res, bookings.ReturnCurrentUserBookings(manager.GetCurrentUser(req)))
	}
}

func serviceHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("Service Handler")

	switch req.Method{
	case "GET":
		id := req.FormValue("id")
		if id == ""{
			break
		}
		p, err := loadPage("book.gohtml")
		if err != nil{
			return
		}

		slot, err := strconv.Atoi(id)
		data := bookings.GenerateCurrentBooking(slot,manager.GetCurrentUser(req))

		p.Body.Execute(res,data)
		return
	default:
		p, err := loadPage("available.gohtml")
		if err != nil{
			return
		}
		p.Body.Execute(res,empty{})
		return
	}


	//title := req.URL.Path[len("/service/"):]
	p, _ := loadPage("available.gohtml")
	fmt.Println(bookings.BookingsData)
	err := p.Body.Execute(res, bookings.BookingsData)
	if err != nil{
		p.Body.Execute(res, empty{})
	}

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

	p, _ := loadPage("index.gohtml")
	fmt.Println(bookings.BookingsData)
	err := p.Body.Execute(res, bookings.BookingsData)
	if err != nil{
		p.Body.Execute(res, empty{})
	}
}

func DELETEALLSESSIONS(res http.ResponseWriter, req *http.Request) {
	fmt.Println("DESTROYING ALL SESSIONS")
	io.WriteString(res, "<html><script type=\"text/javascript\">\nalert(\"DESTROYING ALL SESSIONS\");</script>\n</html>")

	for _,v := range manager.Users{
		manager.DeleteSession(v.Username)
	}
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}

func DELETEALLUSERS(res http.ResponseWriter, req *http.Request) {
	fmt.Println("DESTORYING ALL USERS AND FREEING UP BOOKINGS")
	io.WriteString(res, "<html><script type=\"text/javascript\">alert(\"DESTROYING ALL USERS\");</script>\n</html>")

	for _,v := range manager.Users{
		manager.DeleteUser(v.Username)
		bookings.DeleteBookingByUser(v.Username)
	}
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}