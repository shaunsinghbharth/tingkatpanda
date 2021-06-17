package fetcher

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

	var responseObject []models.CombinedItem
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject
}

func GetAllCombinedItem(key string) []models.CombinedItemToEdit{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/getfullitems/", nil)
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

	var responseObject []models.CombinedItemToEdit
	json.Unmarshal(bodyBytes, &responseObject)

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

	return responseObject
}

func GetShops(key string) []models.Shops{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/getshops/", nil)
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
	var responseObject []models.Shops
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject
}

func GetEditShops(key string) []models.ShopsEdit{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/getshops/", nil)
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
	var responseObject []models.ShopsEdit
	json.Unmarshal(bodyBytes, &responseObject)

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

func EditItem(key string, flag string, itemID string, itemName string, itemPrice string, itemTiming string, itemDesc string, itemImg string, itemCategory string, shopID string) []models.UserItems{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/edititems/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("itemID", itemID)
	q.Add("itemName", itemName)
	q.Add("itemPrice", itemPrice)
	q.Add("itemDesc", itemDesc)
	q.Add("itemImg", itemImg)
	q.Add("itemTiming", itemTiming)
	q.Add("itemCategory", itemCategory)
	q.Add("shopID", shopID)


	req.URL.RawQuery = q.Encode()
	fmt.Println("QUERY: ", req.URL.RawQuery)

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

func DeleteItem(key string, itemID string) []models.UserItems{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/deleteitems/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("itemID", itemID)
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

//("KEYVALUE", flag, shopID, shopName, shopAddress, shopRating, shopPostCode)
func EditShop(key string, flag string, shopID string, shopName string, shopAddress string, shopRating string, shopPostCode string) []models.Shops{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/editshops/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("FETCHEDITSHOP")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("shopID", shopID)
	q.Add("shopName", shopName)
	q.Add("shopAddress", shopAddress)
	q.Add("shopRating", shopRating)
	q.Add("shopPostCode", shopPostCode)

	req.URL.RawQuery = q.Encode()
	fmt.Println("QUERY: ", req.URL.RawQuery)

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
	var responseObject []models.Shops
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API UserItems as struct %+v\n", responseObject)

	return responseObject
}

func DeleteShop(key string, shopID string) []models.UserItems{
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5555/deleteshops/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("FETCHDELETESHOP")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("shopID", shopID)
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