package dto

import "time"

type Order struct {
	ID           string       `json:"id"`
	OrderDate    time.Time    `json:"order_date"`
	ShipDate     time.Time    `json:"ship_date"`
	PostalCode   int          `json:"postal_code"`
	Customer     Customer     `json:"customer"`
	ShipMode     ShipMode     `json:"ship_mode"`
	City         City         `json:"city"`
	Region       Region       `json:"region"`
	OrderDetails OrderDetails `json:"order_details"`
}

type OrderResponse struct {
	Order    Order    `json:"order"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	CurrentPage int `json:"current_page"`
	Pages       int `json:"pages"`
	Total       int `json:"total"`
}

type OrderDetail struct {
	ID       int     `json:"id"`
	Sales    float64 `json:"sales"`
	Quantity int     `json:"quantity"`
	Discount float64 `json:"discount"`
	Profit   float64 `json:"profit"`
	Product  Product `json:"product"`
}

type OrderDetails []OrderDetail

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	SubCategory SubCategory `json:"sub_category"`
}
type Products []Product

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Categories []Category

type SubCategory struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

type SubCategories []SubCategory

type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Regions []Region
type Orders []Order

type Segment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Segments []Segment

type Customer struct {
	ID      string  `json:"json"`
	Name    string  `json:"name"`
	Segment Segment `json:"segment"`
}

type Customers []Customer

type ShipMode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ShipModes []ShipMode

type City struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State State  `json:"state"`
}

type Cities []City

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Countries []Country

type State struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country Country `json:"country"`
}

type States []State

type ProductSummary struct {
	ProductID   string  `json:"id"`
	ProductName string  `json:"product_name"`
	SubCategory string  `json:"sub_category"`
	Category    string  `json:"category"`
	TotalSales  float64 `json:"total_sales"`
}

type ProductSummariesResponse struct {
	Metadata Metadata         `json:"metadata"`
	Data     ProductSummaries `json:"data"`
}

type ProductSummaries []ProductSummary
