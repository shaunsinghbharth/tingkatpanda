package main

import (

	"fmt"
	"tingkatpanda/enginator"
)

func main() {
	// Accessing a new regommend table for the first time will create it.
	western := enginator.Table("western")

	FoodChrisAte := make(map[interface{}]float64)
	FoodChrisAte["Burger"] = 5.0
	FoodChrisAte["Fries"] = 4.0
	FoodChrisAte["Fish"] = 3.0
	western.Add("Chris", FoodChrisAte)

	FoodJayAte := make(map[interface{}]float64)
	FoodJayAte["Burger"] = 3.0
	FoodJayAte["Fries"] = 2.0
	FoodJayAte["Spaghetti"] = 1.5
	western.Add("Jay", FoodJayAte)

	nbs, _ := western.Neighbors("Chris")
	for _, nb := range nbs {
		fmt.Println("Recommending", nb.Key, "with score:", nb.Distance)
	}

	recs, _ := western.Recommend("Chris")
	for _, rec := range recs {
		fmt.Println("Recommending", rec.Key, "with score:", rec.Distance)
	}
}
