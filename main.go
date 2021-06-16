package main

import (
	"sync"
	"tingkatpanda/crud"
	"tingkatpanda/httpd"
)

func init(){

}

func main(){
	var wg sync.WaitGroup

	wg.Add(2)
	go crud.Initialise(wg)
	go httpd.Start(wg)

	wg.Wait()

	// httpd.Stop()
}
