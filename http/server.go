package http

import (
	"assignment3/sessionManager"
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  *template.Template
}


type empty struct{

}

var manager *sessionManager.SessionManager
var srv *http.Server

func Start(){
	fmt.Println("Starting HTTP Server")
	srv = &http.Server{Addr: ":5221"}

	fmt.Println("Initialising Session Manager")
	manager = &sessionManager.SessionManager{}
	manager.Init()

	fs := http.FileServer(http.Dir("templates/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	is := http.FileServer(http.Dir("templates/images/"))
	http.Handle("/images/", http.StripPrefix("/images/", is))

	http.HandleFunc("/", serveHTTP)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/signup/", signupHandler)
	http.HandleFunc("/service/", serviceHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", viewHandler)
	http.HandleFunc("/profile/", profileEditor)
	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/destroy/", Destroy)

	//ADMIN ONLY
	http.HandleFunc("/deletesessions/", DELETEALLSESSIONS)
	http.HandleFunc("/deleteusers/", DELETEALLUSERS)


	//http.HandleFunc("/css/", CSSHandler)

	//http.Handle("/css/", http.StripPrefix("/css/", fileServer))

	err := srv.ListenAndServe()
	if err != nil{
		panic(nil)
		Stop()
	}
	defer srv.Shutdown(nil)
}

func Stop(){
	fmt.Println("Stopping HTTP Server")
	srv.Shutdown(nil)
}

func loadPage(title string) (*Page, error) {
	var filename = "templates/" + title
	body, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}