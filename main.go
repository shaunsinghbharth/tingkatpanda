package main

import (
	"runtime"
	"sync"
	"tingkatpanda/admin"
	"tingkatpanda/crud"
	"tingkatpanda/httpd"
	"tingkatpanda/loadbalancer"
)

func init(){

}

func main(){
	runtime.GOMAXPROCS(1000)

	var wg sync.WaitGroup

	wg.Add(4)
	go crud.Initialise(wg)
	go httpd.Start(wg, "18081")
	//go httpd.Start(wg, "18088")
	go admin.Start(wg, "8085")
	go loadbalancer.Start(wg)

	wg.Wait()
}
