package models

type User struct {
	UserName string
	Password string
}

type Shop struct { // map this type to the record in the Shops table
	ShopId      int
	ShopName    string
	ShopAddress string
	ShopRating  string
	ShopPeriod  string
}

type Item struct { // map this type to the record in the Shops table
	ItemId    int
	ItemName  string
	ItemPrice float64
	ItemDesc  string
	ItemImg   string
	ShopId    int
}
