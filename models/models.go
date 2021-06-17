package models

type Users struct { // map this type to the record in the Users table
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Shops struct { // map this type to the record in the Shops table
	ShopId      int    `json:"shopid"`
	ShopName    string `json:"shopname"`
	ShopAddress string `json:"shopaddress"`
	ShopRating  string `json:"shoprating"`
	ShopStart  string `json:"shopstart"`
	ShopEnd  string `json:"shopend"`
	ShopPostCode  string `json:"shoppostalcode"`
}

type Items struct { // map this type to the record in the Items table
	ItemId    int     `json:"itemid"`
	ItemName  string  `json:"itemname"`
	ItemPrice float64 `json:"itemprice"`
	ItemDesc  string  `json:"itemdesc"`
	ItemImg   string  `json:"itemimg"`
	ItemCategory string `json:"itemcategory"`
	ItemTiming	int	`json:"itemtiming""`
	ShopId    int     `json:"shopid"`
}


type CombinedItem struct {
	ItemId    int     `json:"ItemId"`
	ItemName  string  `json:"ItemName"`
	ItemPrice string `json:"ItemPrice"`
	ItemDesc  string  `json:"ItemDesc"`
	ItemImg   string  `json:"ItemImg"`
	ItemCategory string `json:"ItemCategory"`
	ItemTiming	string	`json:"ItemTiming"`
	ShopId    int     `json:"shopid"`
	ShopName    string `json:"shopname"`
	ShopAddress string `json:"shopaddress"`
	ShopRating  string `json:"shoprating"`
	ShopStart  string `json:"shopstart"`
	ShopEnd  string `json:"shopend"`
	ShopPostCode  string `json:"shoppostalcode"`
}

type UserItems struct {
	ItemID   int    `json:"itemid"`
	Rating   int `json:"rating"`
	UserName string `json:"username"`
}
