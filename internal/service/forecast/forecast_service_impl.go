package forecastService

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"forecasting-be/internal/dto"
	productRepository "forecasting-be/internal/repository/products"
	"net/http"
	"os"
)

type forecastService struct {
	pr productRepository.ProductRepository
}

func NewForecastService(pr productRepository.ProductRepository) forecastService {
	return forecastService{
		pr: pr,
	}
}

func (fs forecastService) ForecastSales(ctx context.Context, month int, year int, subCategoryId int) (dto.ForecastResult, error) {

	if month <= 0 || month > 12 {
		return dto.ForecastResult{}, errors.New("out of boundary")
	}

	data, err := fs.pr.GetSubCategoryByID(ctx, subCategoryId)
	if err != nil {
		return dto.ForecastResult{}, err
	}

	body := fmt.Sprintf(`{
	"month":%v,
	"year":%v,
	"sub_category":"%v"
	}
	`, month, year, data.Name)
	forecastAPIURL := os.Getenv("FORECAST_API_ADDRESS")
	jsonBody := []byte(body)
	res, err := http.Post(forecastAPIURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil || res.StatusCode >= 300 {
		return dto.ForecastResult{}, err
	}
	result := dto.ForecastResponseFromFlask{}

	err = result.FromJSON(res.Body)
	if err != nil {
		return dto.ForecastResult{}, err
	}

	return result.Data, nil
}
