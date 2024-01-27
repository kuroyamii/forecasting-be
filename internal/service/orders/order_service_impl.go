package orderService

import (
	"context"
	"forecasting-be/internal/dto"
	ordersRepository "forecasting-be/internal/repository/orders"
	"forecasting-be/pkg/utilities"
	"log"
)

type orderService struct {
	or ordersRepository.OrderRepository
}

func NewOrderService(or ordersRepository.OrderRepository) orderService {
	return orderService{
		or: or,
	}
}

func (os orderService) GetOrders(ctx context.Context, page int64) (dto.Orders, error) {
	orders, err := os.or.GetOrders(ctx)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return dto.Orders{}, err
	}

	orderDetails, err := os.or.GetOrderDetails(ctx)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return dto.Orders{}, err
	}

	for _, od := range orderDetails {
		ord := orders[od.OrderID]
		ord.OrderDetails = append(ord.OrderDetails, od)
		orders[od.OrderID] = ord
	}

	result := dto.Orders{}
	for _, ord := range orders {
		data := dto.Order{
			ID:         ord.ID,
			OrderDate:  ord.OrderDate,
			ShipDate:   ord.ShipDate,
			PostalCode: ord.PostalCode,
			Customer: dto.Customer{
				ID:      ord.Customer.ID,
				Name:    ord.Customer.Name,
				Segment: dto.Segment(ord.Customer.Segment),
			},
			ShipMode: dto.ShipMode{
				ID:   ord.ShipMode.ID,
				Name: ord.ShipMode.Name,
			},
			City: dto.City{
				ID:   ord.City.ID,
				Name: ord.City.Name,
				State: dto.State{
					ID:      ord.City.State.ID,
					Name:    ord.City.State.Name,
					Country: dto.Country(ord.City.State.Country),
				},
			},
			Region: dto.Region(ord.Region),
		}
		for _, od := range ord.OrderDetails {
			var odt dto.OrderDetail
			odt = dto.OrderDetail{
				ID:       od.ID,
				Sales:    od.Sales,
				Quantity: od.Quantity,
				Discount: od.Discount,
				Profit:   od.Profit,
				Product: dto.Product{
					ID:   od.Product.ID,
					Name: od.Product.Name,
					SubCategory: dto.SubCategory{
						ID:       od.Product.SubCategory.ID,
						Name:     od.Product.SubCategory.Name,
						Category: dto.Category(od.Product.SubCategory.Category),
					},
				},
			}
			data.OrderDetails = append(data.OrderDetails, odt)
		}
		result = append(result, data)
	}

	return result, nil
}

func (os orderService) GetSalesSum(ctx context.Context, month int) (dto.SalesSums, error) {
	res, err := os.or.GetSumsOfSales(ctx, month)
	if err != nil {
		return nil, err
	}
	var result dto.SalesSums
	for _, item := range res {
		data := dto.SalesSum{
			Sum:   item.Sums,
			Month: item.Month,
			Year:  item.Year,
		}
		result = append(result, data)
	}
	return result, err
}

func (os orderService) GetTotalProduct(ctx context.Context) (int64, error) {
	res, err := os.or.GetTotalProducts(ctx)
	return res, err
}

func (os orderService) GetMostBoughtCategory(ctx context.Context) (string, error) {
	res, err := os.or.GetMostBoughtCategory(ctx)
	return res, err
}

func (os orderService) GetTopTransactions(ctx context.Context, limit int) (dto.TopTransactions, error) {
	data, err := os.or.GetTopTransactions(ctx, limit)
	if err != nil {
		return nil, err
	}

	var res dto.TopTransactions
	for _, item := range data {
		transaction := dto.TopTransaction{
			CustomerID: item.CustomerID,
			ItemName:   item.ItemName,
			Date:       item.Date,
			Sales:      item.Sales,
		}
		res = append(res, transaction)
	}
	return res, err
}
