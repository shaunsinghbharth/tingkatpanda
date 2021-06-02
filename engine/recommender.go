package engine

import (
	"fmt"
	"os"
)

var rr Redrec
var err error

func chekErrorAndExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		rr.CloseConn()
		os.Exit(1)
	}
}

func rate(user string, item string, score float64) {
	fmt.Printf("User %s ranked item %s with %.2f\n", user, item, score)
	err := rr.Rate(item, user, score)
	chekErrorAndExit(err)
}

func getProbability(user string, item string) {
	score, err := rr.CalcItemProbability(item, user)
	chekErrorAndExit(err)
	fmt.Printf("%s %s %.2f\n", user, item, score)
}

func suggest(user string, max int) {
	fmt.Printf("Getting %d results for user %s\n", max, user)
	rr.UpdateSuggestedItems(user, max)
	s, err := rr.GetUserSuggestions(user, max)
	chekErrorAndExit(err)
	fmt.Println("results:")
	fmt.Println(s)
}

func update(max int) {
	fmt.Printf("Updating DB\n")
	err := rr.BatchUpdateSimilarUsers(max)
	chekErrorAndExit(err)
}
