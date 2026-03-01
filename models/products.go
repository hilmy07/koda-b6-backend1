package models

type Product struct {
	ID            string  `json:"id"`
	Picture       string  `json:"picture"`
	Name          string  `json:"name"`
	Desc          string  `json:"desc"`
	Price         float64 `json:"price"`
	FlashSale     bool    `json:"flash_sale"`
	Ratings       float64 `json:"ratings"`
	Discount      float64 `json:"discount"`
	ReviewCounter int     `json:"review_counter"`
	Quantity      int     `json:"quantity"`
	Size          string  `json:"size"`
	Variant       string  `json:"variant"`
}

var ListProduct []Product
var ProductCounter int