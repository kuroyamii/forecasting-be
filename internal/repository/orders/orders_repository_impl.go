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
func (or orderRepository) GetSumsOfSales(ctx context.Context, month int) (model.SalesSums, error) {
	rows, err := or.db.QueryContext(ctx, query.GET_SALES_SUM, month)
	if err != nil {
		log.Printf("%v -> %v\n", utilities.Red("ERROR"), err.Error())
		return model.SalesSums{}, err
	}

	var result model.SalesSums
	for rows.Next() {
		var row model.SalesSum
		err = rows.Scan(&row.Sums, &row.Month, &row.Year)
		if err != nil {
			log.Printf("%v -> %v\n", utilities.Red("ERROR"), err.Error())
			return model.SalesSums{}, err
		}
		result = append(result, row)
	}

	return result, err
}
func (or orderRepository) GetTotalProducts(ctx context.Context) (int64, error) {
	rows, err := or.db.QueryContext(ctx, query.GET_TOTAL_PRODUCT)
	if err != nil {
		return 0, err
	}
	var result int64
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return 0, err
		}
	}
	return result, nil
}
func (or orderRepository) GetMostBoughtCategory(ctx context.Context) (string, error) {
	rows, err := or.db.QueryContext(ctx, query.GET_MOST_BOUGHT_CATEGORY)
	if err != nil {
		return "", err
	}

	var result string
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return "", err
		}
	}
	return result, err
}
func (or orderRepository) GetTopTransactions(ctx context.Context, limit int) (model.TopTransactions, error) {
	rows, err := or.db.QueryContext(ctx, query.GET_TOP_TRANSACTION, limit)
	if err != nil {
		return nil, err
	}
	var transactions model.TopTransactions
	for rows.Next() {
		var transaction model.TopTransaction
		err = rows.Scan(&transaction.CustomerID, &transaction.ItemName, &transaction.Date, &transaction.Sales)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
func (or orderRepository) GetProductSummary(ctx context.Context) (model.Products, error) {
	return nil, nil
}

func NewOrderRepository(db *sql.DB) orderRepository {
	return orderRepository{
		db: db,
	}
}
