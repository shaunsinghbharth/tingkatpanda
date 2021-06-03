package main

import (
	"tingkatpanda/engine"
)

func init(){

}

func main(){
	rConn, err := engine.New("redis://redis-16571.c1.us-east1-2.gce.cloud.redislabs.com:16571", "vaqcck1MYIFXbzah3KgJTZN4PvcqHR3m")
	engine.SetConnection(rConn)

	if err != nil{
		//
	}else{
		engine.Rate("user1","food1",5)
		engine.Rate("user1","food6",2)
		engine.Rate("user1","food2",3)
		engine.Rate("user1","food3",4)
		engine.Rate("user1","food5",5)
		engine.Suggest("user1",5)
	}
}