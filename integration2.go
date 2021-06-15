package main

import (
	"sync"
	"tingkatpanda/CRUD"
	"tingkatpanda/httpd"
)

func init(){

}

func main(){
	var wg sync.WaitGroup

	wg.Add(2)
	go CRUD.Initialise(wg)
	go httpd.Start(wg)

	wg.Wait()

	// httpd.Stop()
}
