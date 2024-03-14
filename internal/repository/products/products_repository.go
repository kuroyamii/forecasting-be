package productRepository

import (
	"context"
	"forecasting-be/internal/model"
)

type ProductRepository interface {
	GetProductByID(ctx context.Context, productID int) (model.Product, error)
	GetProductSummary(ctx context.Context, limit int, offset int) (model.ProductSummaries, int, error)
	GetProducts(ctx context.Context) (model.Products, error)
}
