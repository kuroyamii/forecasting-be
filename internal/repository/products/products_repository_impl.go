package productRepository

import (
	"context"
	"database/sql"
	"errors"
	"forecasting-be/internal/model"
	"forecasting-be/internal/query"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) productRepository {
	return productRepository{
		db: db,
	}
}

func (pr productRepository) GetProductByID(ctx context.Context, productID int) (model.Product, error) {
	rows, err := pr.db.QueryContext(ctx, query.GET_PRODUCT_BY_ID, productID)
	if err != nil {
		return model.Product{}, err
	}

	var result model.Product
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.NumericID, &result.SubCategory.ID, &result.SubCategory.Name, &result.SubCategory.Category.ID, &result.SubCategory.Category.Name)
		if err != nil {
			return model.Product{}, err
		}
		return result, nil
	}
	return model.Product{}, errors.New("product not found")
}

func (pr productRepository) GetProductSummary(ctx context.Context, limit int, offset int) (model.ProductSummaries, int, error) {
	rows, err := pr.db.QueryContext(ctx, query.GET_PRODUCT_SUMMARY, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	var total int
	var results model.ProductSummaries

	if !rows.Next() {
		return nil, 0, errors.New("beyond the boundary")
	}
	for rows.Next() {
		var result model.ProductSummary
		err = rows.Scan(&total, &result.ProductID, &result.ProductName, &result.SubCategory, &result.Category, &result.TotalSales)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, result)
	}
	return results, total, nil
}
