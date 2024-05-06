package dto

import (
	"encoding/json"
	"io"
)

type RequestForecast struct {
	ProductID string `json:"product_id"`
	Month     string `json:"month"`
	Year      string `json:"year"`
}

type ResponseForecast struct {
	ProductName string  `json:"product_name"`
	Result      float64 `json:"result"`
}

type ForecastRequest struct {
	Month         int `json:"month" validate:"required"`
	Year          int `json:"year" validate:"required"`
	SubCategoryID int `json:"sub_category_id" validate:"required"`
}

type ForecastResponseFromFlask struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Errors  interface{}    `json:"errors"`
	Data    ForecastResult `json:"data"`
}

type ForecastResult struct {
	Month       int     `json:"month"`
	Year        int     `json:"year"`
	SubCategory string  `json:"sub_category"`
	Result      float64 `json:"result"`
}

func (fr *ForecastResponseFromFlask) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(fr)
}
