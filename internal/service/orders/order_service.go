package orderService

import (
	"context"
	"forecasting-be/internal/dto"
)

type OrderService interface {
	GetOrders(ctx context.Context, page int64) (dto.Orders, error)
	GetSalesSum(ctx context.Context, month int) (dto.SalesSums, error)
	GetTotalProduct(ctx context.Context) (int64, error)
	GetMostBoughtCategory(ctx context.Context) (string, error)
	GetTopTransactions(ctx context.Context, limit int) (dto.TopTransactions, error)
}
