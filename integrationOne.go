package main

import (
	"fmt"
	"strconv"
	"tingkatpanda/enginator"
	"tingkatpanda/myconnector"
)

func main() {
	db := myconnector.ConnectShops()

	western := enginator.Table("western")

	FoodChrisAte := make(map[interface{}]float64)
	FoodChrisAte[1] = 5.0
	FoodChrisAte[2] = 4.0
	FoodChrisAte[3] = 3.0
	western.Add("Chris", FoodChrisAte)

	FoodJayAte := make(map[interface{}]float64)
	FoodJayAte[1] = 3.0
	FoodJayAte[3] = 2.0
	FoodJayAte[5] = 1.5
	western.Add("Jay", FoodJayAte)

	nbs, _ := western.Neighbors("Chris")
	for _, nb := range nbs {
		value, _ := strconv.Atoi(nb.Key.(string))
		fmt.Println("Recommending", myconnector.GetSpecificItemRecord(&db, value), "with score:", nb.Distance)
	}

	recs, _ := western.Recommend("Chris")
	for _, rec := range recs {
		fmt.Println("Recommending", myconnector.GetSpecificItemRecord(&db, rec.Key.(int)), "with score:", rec.Distance)
	}
}

