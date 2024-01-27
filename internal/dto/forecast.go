package dto

type RequestForecast struct {
	ProductID string `json:"product_id"`
	Month     string `json:"month"`
	Year      string `json:"year"`
}

type ResponseForecast struct {
	ProductName string  `json:"product_name"`
	Result      float64 `json:"result"`
}
