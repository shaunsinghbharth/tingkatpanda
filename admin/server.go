package admin

import (
	"fmt"
	"html/template"
	"net/http"
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

func Start(wg sync.WaitGroup, port string){
	fmt.Println("Generating cache")
	c = cachinator.New(5*time.Minute, 10*time.Minute)

	fmt.Println("Creating MUX")
	mux := http.NewServeMux()
	fmt.Println("Starting HTTP Server")
	srv = &http.Server{Addr: ":" + port,
		Handler: mux}

	fmt.Println("Initialising Session Manager")
	manager = &sessionManager.SessionManager{}
	manager.Init("tingkatpandaadmin")

	fs :=http.FileServer(http.Dir("htdocs/css/"))
	mux.Handle("/css/", http.StripPrefix("/css/", fs))
	is := http.FileServer(http.Dir("htdocs/images/"))
	mux.Handle("/images/", http.StripPrefix("/images/", is))
	rs := http.FileServer(http.Dir("htdocs/reference_files/"))
	mux.Handle("/reference_files/", http.StripPrefix("/reference_files/", rs))


	mux.HandleFunc("/", ServeHTTP)

	mux.HandleFunc("/admin/login/", Login)
	mux.HandleFunc("/admin/authenticate", Authenticate)
	mux.HandleFunc("/admin/", ServeFunctions)
	mux.HandleFunc("/admin/functions/shops/", ServeShops)
	mux.HandleFunc("/admin/functions/items/", ServeItems)
	mux.HandleFunc("/admin/functions/users/", ServeUsers)

	mux.HandleFunc("/destroy/", Destroy)
	err := srv.ListenAndServe()
	if err != nil{
		fmt.Println(err)
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