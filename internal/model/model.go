package model

import "time"

type Order struct {
	ID         string    `db:"id"`
	OrderDate  time.Time `db:"order_date"`
	ShipDate   time.Time `db:"ship_date"`
	PostalCode int       `db:"postal_code"`
	Customer   Customer
	ShipMode   ShipMode
	City       City
	Region     Region
}

type OrderDetail struct {
	ID       int     `db:"id"`
	Sales    float64 `db:"sales"`
	Quantity int     `db:"quantity"`
	Discount float64 `db:"discount"`
	Profit   float64 `db:"profit"`
	Product  Product
}

type Product struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	SubCategory SubCategory
}
type Products []Product

type Category struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
type Categories []Category

type SubCategory struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	CategoryID int    `db:"category_id"`
}

type SubCategories []SubCategory

type Region struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Regions []Region
type Orders []Order

type Segment struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Segments []Segment

type Customer struct {
	ID        string `db:"db"`
	Name      string `db:"name"`
	SegmentID int    `db:"segment_id"`
}

type Customers []Customer

type ShipMode struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type ShipModes []ShipMode

type City struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	StateID int    `db:"state_id"`
}

type Cities []City

type Country struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Countries []Country

type State struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type States []State
