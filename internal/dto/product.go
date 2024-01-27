package dto

type ResponseProductSummary struct {
	ProductID         string  `json:"product_id"`
	ProductName       string  `json:"product_name"`
	TotalTransactions int     `json:"total_transactions"`
	TotalSales        float64 `json:"total_sales"`
}

type ResponseProductSummaries []ResponseProductSummary
