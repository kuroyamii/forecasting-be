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

func (pr productRepository) GetSubCategoryByID(ctx context.Context, subCategoryId int) (model.SubCategory, error) {
	rows, err := pr.db.QueryContext(ctx, query.GET_SUB_CATEGORY_BY_ID, subCategoryId)
	if err != nil {
		return model.SubCategory{}, err
	}

	var result model.SubCategory
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.Category.ID, &result.Category.Name)
		if err != nil {
			return model.SubCategory{}, err
		}
		return result, nil
	}
	return model.SubCategory{}, errors.New("product not found")
}

func (pr productRepository) GetProductSummary(ctx context.Context, limit int, offset int) (model.ProductSummaries, int, error) {
	rows, err := pr.db.QueryContext(ctx, query.GET_PRODUCT_SUMMARY, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	var total int
	var results model.ProductSummaries

	for rows.Next() {
		var result model.ProductSummary
		err = rows.Scan(&total, &result.ProductID, &result.ProductName, &result.SubCategory, &result.Category, &result.TotalSales)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, 0, errors.New("beyond the boundary")
	}
	return results, total, nil
}

func (pr productRepository) GetSubCategories(ctx context.Context) (model.SubCategories, error) {
	rows, err := pr.db.QueryContext(ctx, query.GET_SUB_CATEGORY)
	if err != nil {
		return nil, err
	}
	var result model.SubCategories
	for rows.Next() {
		var item model.SubCategory
		err = rows.Scan(&item.ID, &item.Name)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
