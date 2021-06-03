package enginator

import (
	"testing"
)

var (
)

func TestEngine(t *testing.T) {
	western := Table("western")

	FoodChrisAte := make(map[interface{}]float64)
	FoodChrisAte["Burger"] = 5.0
	FoodChrisAte["Fries"] = 4.0
	FoodChrisAte["Fish"] = 3.0
	western.Add("Chris", FoodChrisAte)

	FoodJayAte := make(map[interface{}]float64)
	FoodJayAte["Burger"] = 5.0
	FoodJayAte["Fries"] = 4.0
	FoodJayAte["Fish"] = 4.5
	western.Add("Jay", FoodJayAte)

	// check if both items are still there
	p, err := western.Value("Chris")
	if err != nil || p == nil {
		t.Error("Error retrieving item from engine", err)
	}
	p, err = western.Value("Jay")
	if err != nil || p == nil {
		t.Error("Error retrieving item from engine", err)
	}
}

func TestNeighbors(t *testing.T) {
	western := Table("western")

	FoodChrisAte := make(map[interface{}]float64)
	FoodChrisAte["Burger"] = 5.0
	FoodChrisAte["Fries"] = 4.0
	FoodChrisAte["Fish"] = 3.0
	western.Add("Chris", FoodChrisAte)

	FoodJayAte := make(map[interface{}]float64)
	FoodJayAte["Burger"] = 5.0
	FoodJayAte["Fries"] = 4.0
	FoodJayAte["Fish"] = 4.5
	western.Add("Jay", FoodJayAte)

	foodMaryAte := make(map[interface{}]float64)
	foodMaryAte["Burger"] = 4.0
	foodMaryAte["Fries"] = 3.0
	foodMaryAte["Fish"] = 4.5
	western.Add("Mary", foodMaryAte)

	foodJackAte := make(map[interface{}]float64)
	foodJackAte["Burger"] = 3.0
	foodJackAte["Fries"] = 1.0
	western.Add("Jack", foodJackAte)

	nbs, _ := western.Neighbors("Chris")
	if len(nbs) != 3 {
		t.Error("Expected 3 neighbours, got", len(nbs))
	}
	if nbs[0].Key != "Jay" || nbs[1].Key != "Mary" || nbs[2].Key != "Jack" {
		t.Error("Unexpected similarity order")
	}
}

func TestRecommendations(t *testing.T) {
	western := Table("western")

	FoodChrisAte := make(map[interface{}]float64)
	FoodChrisAte["Burger"] = 5.0
	FoodChrisAte["Fries"] = 4.0
	FoodChrisAte["Fish"] = 3.0
	western.Add("Chris", FoodChrisAte)

	FoodJayAte := make(map[interface{}]float64)
	FoodJayAte["Western"] = 5.0
	FoodJayAte["Chicken"] = 4.0
	FoodJayAte["Nuggets"] = 4.5
	western.Add("Jay", FoodJayAte)

	recs, _ := western.Recommend("Chris")
	if len(recs) != 2 {
		t.Error("Expected 2 recommendations, got", len(recs))
	}
	if recs[0].Key != "Burger" || recs[1].Key != "Fries" {
		t.Error("Unexpected recommendation order")
	}
}
