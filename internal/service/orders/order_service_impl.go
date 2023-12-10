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

func (os orderService) GetOrders(ctx context.Context) (dto.Orders, error) {
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
		ordDetails := dto.OrderDetails{}
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
			ordDetails = append(ordDetails, odt)
		}
		result = append(result, data)
	}

	return result, nil
}
