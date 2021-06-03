package engine

import (
	"fmt"
	"os"
)

var err error
var rr RedConnection

func SetConnection(redConn *RedConnection){
	rr = *redConn
}

func checkErrorAndExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		rr.CloseConn()
		os.Exit(1)
	}
}

func Rate(user string, item string, score float64) {
	fmt.Printf("User %s ranked item %s with %.2f\n", user, item, score)
	err := rr.Rate(item, user, score)
	checkErrorAndExit(err)
}

func GetProbability(user string, item string) {
	score, err := rr.CalcItemProbability(item, user)
	checkErrorAndExit(err)
	fmt.Printf("%s %s %.2f\n", user, item, score)
}

func Suggest(user string, max int) {
	fmt.Printf("Getting %d results for user %s\n", max, user)
	rr.UpdateSuggestedItems(user, max)
	s, err := rr.GetUserSuggestions(user, max)
	checkErrorAndExit(err)
	fmt.Println("results:")
	fmt.Println(s)
}

func Update(max int) {
	fmt.Printf("Updating DB\n")
	err := rr.BatchUpdateSimilarUsers(max)
	checkErrorAndExit(err)
}