package model

import "time"

type Order struct {
	ID           string    `db:"id"`
	OrderDate    time.Time `db:"order_date"`
	ShipDate     time.Time `db:"ship_date"`
	PostalCode   int       `db:"postal_code"`
	Customer     Customer
	ShipMode     ShipMode
	City         City
	Region       Region
	OrderDetails OrderDetails
}

type OrderDetail struct {
	ID       int     `db:"id"`
	Sales    float64 `db:"sales"`
	Quantity int     `db:"quantity"`
	Discount float64 `db:"discount"`
	Profit   float64 `db:"profit"`
	OrderID  string  `db:"order_id"`
	Product  Product
}

type OrderDetails []OrderDetail

type Product struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	SubCategory SubCategory
}
type Products []Product

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
type Categories []Category

type SubCategory struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Category Category
}

type SubCategories []SubCategory

type Region struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Regions []Region
type Orders map[string]Order

type Segment struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Segments []Segment

type Customer struct {
	ID      string `db:"db"`
	Name    string `db:"name"`
	Segment Segment
}

type Customers []Customer

type ShipMode struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type ShipModes []ShipMode

type City struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	State State
}

type Cities []City

type Country struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Countries []Country

type State struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Country Country
}

type States []State

type SalesSum struct {
	Year  int     `db:"year"`
	Month int     `db:"month"`
	Sums  float64 `db:"sum"`
}

type SalesSums []SalesSum

type TopTransaction struct {
	CustomerID string    `db:"customer_id"`
	ItemName   string    `db:"product_name"`
	Date       time.Time `db:"order_date"`
	Sales      float64   `db:"sales"`
}

type TopTransactions []TopTransaction

type ProductSummary struct {
	ProductID   string  `db:"id"`
	ProductName string  `db:"product_name"`
	SubCategory string  `db:"sub_category"`
	Category    string  `db:"category"`
	TotalSales  float64 `db:"total_sales"`
}

type ProductSummaries []ProductSummary
