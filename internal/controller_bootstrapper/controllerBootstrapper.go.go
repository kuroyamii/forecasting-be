package controllerBootstrapper

import (
	"database/sql"
	orderController "forecasting-be/internal/controller/orders"
	pingController "forecasting-be/internal/controller/ping"
	ordersRepository "forecasting-be/internal/repository/orders"
	orderService "forecasting-be/internal/service/orders"
	pingService "forecasting-be/internal/service/ping"

	"github.com/gorilla/mux"
)

func InitializeEndpoints(router *mux.Router, db *sql.DB) {
	pingService := pingService.NewPingService()

	pingController := pingController.NewPingController(router, pingService)
	pingController.InitEndpoints()

	orderRepository := ordersRepository.NewOrderRepository(db)
	orderService := orderService.NewOrderService(orderRepository)
	orderController := orderController.NewOrderController(router, orderService)
	orderController.InitializeEndpoints()

}
