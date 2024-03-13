package productService

import (
	"context"
	"forecasting-be/internal/dto"
)

type ProductService interface {
	GetProductSummary(ctx context.Context, page int, limit int) (dto.ProductSummariesResponse, error)
}
