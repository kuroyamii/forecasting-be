package ordersRepository

import (
	"context"
	"database/sql"
	"forecasting-be/internal/model"
	"forecasting-be/internal/query"
	"forecasting-be/pkg/utilities"
	"log"
)

type orderRepository struct {
	db *sql.DB
}

func (or orderRepository) GetOrders(ctx context.Context) (model.Orders, error) {
	result, err := or.db.QueryContext(ctx, query.GetOrders)
	if err != nil {
		log.Printf("%v -> %v\n", utilities.Red("ERROR"), err.Error())
		return model.Orders{}, err
	}
	orders := model.Orders{}

	for result.Next() {
		var order model.Order
		// var od model.OrderDetail
		// err = result.Scan(&order.ID, &order.OrderDate, &order.ShipDate, &order.PostalCode, &order.Customer.ID, &order.Customer.Name, &order.Customer.Segment.ID, &order.Customer.Segment.Name, &order.ShipMode.ID, &order.ShipMode.Name, &order.City.ID, &order.City.Name, &order.City.State.ID, &order.City.State.Name, &order.City.State.Country.ID, &order.City.State.Country.Name, &order.Region.ID, &order.Region.Name, &od.ID, &od.Sales, &od.Quantity, &od.Discount, &od.Profit, &od.Product.ID, &od.Product.Name, &od.Product.SubCategory.ID, &od.Product.SubCategory.Name, &od.Product.SubCategory.Category.ID, &od.Product.SubCategory.Category.Name)
		err = result.Scan(&order.ID, &order.OrderDate, &order.ShipDate, &order.PostalCode, &order.Customer.ID, &order.Customer.Name, &order.Customer.Segment.ID, &order.Customer.Segment.Name, &order.ShipMode.ID, &order.ShipMode.Name, &order.City.ID, &order.City.Name, &order.City.State.ID, &order.City.State.Name, &order.City.State.Country.ID, &order.City.State.Country.Name, &order.Region.ID, &order.Region.Name)
		if err != nil {
			log.Printf("%v -> %v\n", utilities.Red("ERROR"), err.Error())
			return model.Orders{}, err
		}
		// order.OrderDetail = od
		// orders = append(orders, order)
		orders[order.ID] = order
	}
	return orders, nil

}

func (or orderRepository) GetOrderDetails(ctx context.Context) (model.OrderDetails, error) {
	result, err := or.db.QueryContext(ctx, query.GetOrderDetails)
	if err != nil {
		log.Printf("%v -> %v\n", utilities.Red("ERROR"), err.Error())
		return model.OrderDetails{}, err
	}
	orderDetails := model.OrderDetails{}

	for result.Next() {
		var od model.OrderDetail
		err = result.Scan(&od.ID, &od.Sales, &od.Quantity, &od.Discount, &od.Profit, &od.OrderID, &od.Product.ID, &od.Product.Name, &od.Product.SubCategory.ID, &od.Product.SubCategory.Name, &od.Product.SubCategory.Category.ID, &od.Product.SubCategory.Category.Name)
		if err != nil {
			log.Printf("%v -> %v\n", utilities.Red("ERROR"), err.Error())
			return model.OrderDetails{}, err
		}
		orderDetails = append(orderDetails, od)

	}
	return orderDetails, nil
}

func NewOrderRepository(db *sql.DB) orderRepository {
	return orderRepository{
		db: db,
	}
}
