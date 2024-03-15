package productRepository

import (
	"context"
	"forecasting-be/internal/model"
)

type ProductRepository interface {
	GetSubCategoryByID(ctx context.Context, subCategoryId int) (model.SubCategory, error)
	GetProductSummary(ctx context.Context, limit int, offset int) (model.ProductSummaries, int, error)
	GetSubCategories(ctx context.Context) (model.SubCategories, error)
}
