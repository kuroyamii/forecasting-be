package productService

import (
	"context"
	"forecasting-be/internal/dto"
	productRepository "forecasting-be/internal/repository/products"
)

type productService struct {
	pr productRepository.ProductRepository
}

func NewProductService(pr productRepository.ProductRepository) productService {
	return productService{
		pr: pr,
	}
}

func (ps productService) GetProductSummary(ctx context.Context, page int, limit int) (dto.ProductSummariesResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := limit*(page-1) + 1
	data, total, err := ps.pr.GetProductSummary(ctx, limit, offset)
	totalPages := total / limit
	if err != nil {
		return dto.ProductSummariesResponse{}, err
	}
	var result dto.ProductSummaries
	for _, item := range data {
		res := dto.ProductSummary{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Category:    item.Category,
			SubCategory: item.SubCategory,
			TotalSales:  item.TotalSales,
		}
		result = append(result, res)
	}

	response := dto.ProductSummariesResponse{
		Metadata: dto.Metadata{
			CurrentPage: page,
			Pages:       totalPages,
			Total:       total,
		},
		Data: result,
	}
	return response, nil
}
