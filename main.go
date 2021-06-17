package main

import (
	"sync"
	"tingkatpanda/admin"
	"tingkatpanda/crud"
	"tingkatpanda/httpd"
	"tingkatpanda/loadbalancer"
)

func init(){

}

func main(){
	var wg sync.WaitGroup

	wg.Add(4)
	go crud.Initialise(wg)
	go httpd.Start(wg, "8081")
	go admin.Start(wg, "8085")
	go loadbalancer.Start(wg)

	wg.Wait()
}
