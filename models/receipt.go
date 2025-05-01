package models

// Receipt details
type Receipt struct {
	Retailer    string `json:"retailer"`
	Date        string `json:"purchaseDate"`
	Time        string `json:"purchaseTime"`
	Items       []Item `json:"items"`
	TotalAmount string `json:"total"`
}

// Item details
type Item struct {
	Description string `json:"shortDescription"`
	Price       string `json:"price"`
}
