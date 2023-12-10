package ordersRepository

import (
	"context"
	"forecasting-be/internal/model"
)

type OrderRepository interface {
	GetOrders(ctx context.Context) (model.Orders, error)
	GetOrderDetails(ctx context.Context) (model.OrderDetails, error)
}
