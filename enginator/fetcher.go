package enginator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tingkatpanda/models"
)

func GetCombinedItem(key string, itemID string) []models.CombinedItem{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/getfullitem/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("item", itemID)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(string(bodyBytes))
	var responseObject []models.CombinedItem
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Course as struct %+v\n", responseObject)

	fmt.Println("COMBINED ITEMS ", responseObject)
	return responseObject
}

func GetUsers(key string) []models.Users{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/getusers/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(string(bodyBytes))
	var responseObject []models.Users
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Course as struct %+v\n", responseObject)

	return responseObject
}

func GetUser(key string, userID string) []models.Users{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/getspecificusers/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("user", userID)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(string(bodyBytes))
	var responseObject []models.Users
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Course as struct %+v\n", responseObject)

	return responseObject
}

func GetUserItems(key string, userID string) []models.UserItems{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/getuseritems/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("user", userID)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(string(bodyBytes))
	var responseObject []models.UserItems
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API UserItems as struct %+v\n", responseObject)

	return responseObject
}