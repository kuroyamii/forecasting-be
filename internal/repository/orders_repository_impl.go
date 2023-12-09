package ordersRepository

import (
	"database/sql"
)

type orderRepository struct {
	db *sql.DB
}

func (or orderRepository) GetOrders() {

}
