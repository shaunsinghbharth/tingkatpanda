package httpd

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
	"time"
	"tingkatpanda/cachinator"
	"tingkatpanda/sessionManager"
)

type Page struct {
	Title string
	Body  *template.Template
}

type empty struct{

}

var manager *sessionManager.SessionManager
var srv *http.Server
var c *cachinator.Cache

func Start(wg sync.WaitGroup){
	fmt.Println("Generating cache")
	c = cachinator.New(5*time.Minute, 10*time.Minute)

	fmt.Println("Starting HTTP Server")
	srv = &http.Server{Addr: ":" + os.Getenv("PORT")}

	fmt.Println("Initialising Session Manager")
	manager = &sessionManager.SessionManager{}
	manager.Init("tingkatpanda")

	fs := http.FileServer(http.Dir("htdocs/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	is := http.FileServer(http.Dir("htdocs/images/"))
	http.Handle("/images/", http.StripPrefix("/images/", is))
	rs := http.FileServer(http.Dir("htdocs/reference_files/"))
	http.Handle("/reference_files/", http.StripPrefix("/reference_files/", rs))


	http.HandleFunc("/", ServeHTTP)
	http.HandleFunc("/recommend/", ShowRecommendation)
	http.HandleFunc("/select/", ShowSelect)
	http.HandleFunc("/login/", Login)
	http.HandleFunc("/authenticate/", Authenticate)

	http.HandleFunc("/destroy/", Destroy)
	err := srv.ListenAndServe()
	if err != nil{
		panic(nil)
		Stop()
	}
	defer srv.Shutdown(nil)
	wg.Done()
}

func Stop(){
	fmt.Println("Stopping HTTP Server")
	srv.Shutdown(nil)
}

func loadPage(title string) (*Page, error) {
	var filename = "htdocs/" + title
	body, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}