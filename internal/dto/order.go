package dto

import "time"

type SalesSum struct {
	Sum   float64 `json:"sum"`
	Month int     `json:"month"`
	Year  int     `json:"year"`
}

type SalesSums []SalesSum

type TopTransaction struct {
	CustomerID string    `json:"customer_id"`
	ItemName   string    `json:"product_name"`
	Date       time.Time `json:"order_date"`
	Sales      float64   `json:"sales"`
}

type TopTransactions []TopTransaction
