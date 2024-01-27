package ordersRepository

import (
	"context"
	"forecasting-be/internal/model"
)

type OrderRepository interface {
	GetOrders(ctx context.Context) (model.Orders, error)
	GetOrderDetails(ctx context.Context) (model.OrderDetails, error)
	GetSumsOfSales(ctx context.Context, month int) (model.SalesSums, error)
	GetTotalProducts(ctx context.Context) (int64, error)
	GetMostBoughtCategory(ctx context.Context) (string, error)
	GetTopTransactions(ctx context.Context, limit int) (model.TopTransactions, error)
	GetProductSummary(ctx context.Context) (model.Products, error)
}
