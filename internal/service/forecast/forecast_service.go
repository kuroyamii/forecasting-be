package forecastService

import (
	"context"
	"forecasting-be/internal/dto"
)

type ForecastServie interface {
	ForecastSales(ctx context.Context, month int, year int, subCategoryId int) (dto.ForecastResult, error)
}
