package orderService

import (
	"context"
	"forecasting-be/internal/dto"
)

type OrderService interface {
	GetOrders(ctx context.Context, page int64) (dto.Orders, error)
}
